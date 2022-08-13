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
