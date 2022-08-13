package truck

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type truckRepository struct {
	dbName string
}

var (
	sortMap = map[string]string{
		"truckType":       "truck_types.name ASC",
		"-truckType":      "truck_types.name DESC",
		"licenseNumber":   "license_number ASC",
		"-licenseNumber":  "license_number DESC",
		"plateType":       "plate_type ASC",
		"-plateType":      "plate_type DESC",
		"productionYear":  "production_year ASC",
		"-productionYear": "production_year DESC",
		"status":          "status ASC",
		"-status":         "status DESC",
		"stnk":            "stnk ASC",
		"-stnk":           "stnk DESC",
		"kir":             "kir ASC",
		"-kir":            "kir DESC",
	}
)

func NewRepository(dbName string) TruckRepositoryInterface {
	return &truckRepository{
		dbName: dbName,
	}
}

func (repo *truckRepository) selectFields() []string {
	return []string{
		"trucks.id",
		"truck_types.name AS truck_type",
		"license_number",
		"plate_type",
		"production_year",
		"status",
		"stnk",
		"kir",
	}
}

func (repo *truckRepository) GetTableName() string {
	return "trucks"
}

func (repo *truckRepository) GetTrucks(ctx context.Context, db *sqlx.DB, filter *models.TruckFilter) ([]models.Truck, error) {
	var result []models.Truck

	builder := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Join("truck_types ON trucks.truck_type_id = truck_types.id").
		Where(sq.Eq{"status": true})

	if filter != nil {
		if filter.Search != "" {
			builder = builder.Where(sq.Or{
				sq.Like{"license_number": "%" + filter.Search + "%"},
			})
		}

		if filter.TruckTypeID != 0 {
			builder = builder.Where(sq.Eq{
				"truck_types.id": filter.TruckTypeID,
			})
		}

		sortBy := sortMap[filter.Sort]
		if sortBy == "" {
			sortBy = "license_number DESC"
		}
		builder = builder.OrderBy(sortBy)
	}

	query, args, _ := builder.ToSql()

	err := db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTrucks"}).
			Str("query", query).
			Msg(err.Error())

		return result, err
	}

	return result, nil
}

func (repo *truckRepository) GetTruckByID(ctx context.Context, db *sqlx.DB, truckID int64) (*models.Truck, error) {
	query, args, _ := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Join("truck_types ON trucks.truck_type_id = truck_types.id").
		Where(sq.Eq{"trucks.id": truckID}).
		ToSql()

	var result models.Truck
	err := db.GetContext(ctx, &result, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTruckByID"}).
			Str("query", query).
			Msg(err.Error())

		return nil, err
	}

	return &result, nil
}

func (repo *truckRepository) GetTruckByLicenseNumber(ctx context.Context, db *sqlx.DB, licenseNumber string, truckID *int64) (*models.Truck, error) {
	builder := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Join("truck_types ON trucks.truck_type_id = truck_types.id").
		Where(sq.Eq{"trucks.license_number": licenseNumber})

	if truckID != nil {
		builder = builder.Where(sq.NotEq{"trucks.id": truckID})
	}

	query, args, _ := builder.ToSql()

	var result models.Truck
	err := db.GetContext(ctx, &result, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTruckByLicenseNumber"}).
			Str("query", query).
			Msg(err.Error())

		return nil, err
	}

	return &result, nil
}

func (repo *truckRepository) Create(ctx context.Context, db *sqlx.DB, payload *models.CreateTruckPayload) error {
	insertMap := sq.Eq{
		"truck_type_id":   payload.TruckTypeID,
		"license_number":  payload.LicenseNumber,
		"plate_type":      payload.PlateType,
		"production_year": payload.ProductionYear,
		"status":          1,
	}

	if payload.STNKUrl != "" {
		insertMap["stnk"] = payload.STNKUrl
	}

	if payload.KirUrl != "" {
		insertMap["kir"] = payload.KirUrl
	}

	query, args, _ := sq.Insert(repo.GetTableName()).
		SetMap(insertMap).
		ToSql()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "Create"}).
			Str("query", query).
			Msg(err.Error())

		return err
	}

	return nil
}

func (repo *truckRepository) Update(ctx context.Context, db *sqlx.DB, payload *models.CreateTruckPayload, truckID int64) error {
	updateMap := sq.Eq{
		"truck_type_id":   payload.TruckTypeID,
		"license_number":  payload.LicenseNumber,
		"plate_type":      payload.PlateType,
		"production_year": payload.ProductionYear,
	}

	if payload.STNKUrl != "" {
		updateMap["stnk"] = payload.STNKUrl
	}

	query, args, _ := sq.Update(repo.GetTableName()).
		SetMap(updateMap).
		Where(sq.Eq{"id": truckID}).
		ToSql()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "Update"}).
			Str("query", query).
			Msg(err.Error())

		return err
	}

	return nil
}

func (repo *truckRepository) DeactivateTruck(ctx context.Context, db *sqlx.DB, truckID int64) error {
	query, args, _ := sq.Update(repo.GetTableName()).
		SetMap(sq.Eq{"status": false}).
		Where(sq.Eq{"id": truckID}).
		ToSql()

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "DeactivateTruck"}).
			Str("query", query).
			Msg(err.Error())

		return err
	}

	return nil
}
