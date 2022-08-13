package shipment

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type shipmentRepository struct {
	dbName string
}

func NewRepository(dbName string) ShipmentRepositoryInterface {
	return &shipmentRepository{
		dbName: dbName,
	}
}

var (
	sortMap = map[string]string{
		"shipmentNumber":       "shipment_number ASC",
		"-shipmentNumber":      "shipment_number DESC",
		"licenseNumber":        "license_number ASC",
		"-licenseNumber":       "license_number DESC",
		"driverName":           "drivers.name ASC",
		"-driverName":          "drivers.name DESC",
		"originDistrict":       "o.name ASC",
		"-originDistrict":      "o.name DESC",
		"destinationDistrict":  "d.name ASC",
		"-destinationDistrict": "d.name DESC",
	}
)

func (repo *shipmentRepository) GetTableName() string {
	return "shipments"
}

func (repo *shipmentRepository) selectFields() []string {
	return []string{
		"id",
		"shipment_number",
		"license_number",
		"drivers.name AS driver_name",
		"o.name AS origin_district",
		"d.name AS destination_district",
	}
}

func (repo *shipmentRepository) GetShipments(ctx context.Context, db *sqlx.DB, filter *models.ShipmentFilter) ([]models.Shipment, error) {
	var result []models.Shipment

	builder := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Join("drivers ON shipments.driver_id = drivers.id").
		Join("districts o ON shipments.origin_id = o.id").
		Join("districts d ON shipments.destination_id = d.id").
		Where(sq.Eq{"status": true})

	if filter.Search != "" {
		builder = builder.Where(sq.Like{"shipment_number": "%" + filter.Search + "%"})
	}

	sortBy := sortMap[filter.Sort]
	if sortBy == "" {
		sortBy = "shipment_number DESC"
	}
	builder = builder.OrderBy(sortBy)

	query, args, _ := builder.ToSql()

	err := db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (repo *shipmentRepository) CreateShipment(ctx context.Context, db *sqlx.DB, payload *models.CreateShipmentPayload) error {
	query, args, _ := sq.Insert(repo.GetTableName()).
		SetMap(sq.Eq{
			"shipment_number": payload.ShipmentNumber,
			"origin_id":       payload.OriginID,
			"destination_id":  payload.DestinationID,
			"loading_date":    payload.LoadingDate,
			"status":          true,
		}).ToSql()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *shipmentRepository) AllocateShipment(ctx context.Context, db *sqlx.DB, payload *models.AllocateShipmentPayload, shipmentId int64) error {
	query, args, _ := sq.Update(repo.GetTableName()).
		SetMap(sq.Eq{
			"driver_id": payload.DriverID,
			"truck_id":  payload.TruckID,
		}).Where(sq.Eq{"id": shipmentId}).
		ToSql()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"ShipmentRepository", "AllocateShipment"}).
			Str("query", query).
			Msg(err.Error())

		return err
	}

	return nil
}

func (repo *shipmentRepository) UpdateStatus(ctx context.Context, db *sqlx.DB, payload *models.UpdateShipmentStatusPayload, shipmentId int64) error {
	query, args, _ := sq.Update(repo.GetTableName()).
		SetMap(sq.Eq{"status": payload.Status}).
		Where(sq.Eq{"id": shipmentId}).
		ToSql()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"ShipmentRepository", "UpdateStatus"}).
			Str("query", query).
			Msg(err.Error())

		return err
	}

	return nil
}
