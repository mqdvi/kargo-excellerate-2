package driver

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

var (
	sortMap = map[string]string{
		"name":         "name ASC",
		"-name":        "name DESC",
		"phoneNumber":  "phone_number ASC",
		"-phoneNumber": "phone_number DESC",
		"createdAt":    "created_at ASC",
		"-createdAt":   "created_at DESC",
		"status":       "status ASC",
		"-status":      "status DESC",
	}
)

func (repo *driverRepository) GetDriver(ctx context.Context, db *sqlx.DB, payload *models.GetDriverFilter) ([]models.Driver, error) {
	var result []models.Driver

	builder := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Where(sq.Eq{"status": true})

	if payload.Search != "" {
		builder = builder.Where(sq.Eq{"name": "%" + payload.Search + "%"})
	}

	sortBy := sortMap[payload.Sort]
	if sortBy == "" {
		sortBy = "id DESC"
	}
	builder = builder.OrderBy(sortBy)

	query, args, _ := builder.ToSql()

	err := db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"DriverRepository", "GetDriver"}).
			Str("query", query).
			Msg(err.Error())
		return result, err
	}
	return result, nil
}

func (repo *driverRepository) GetDriverByID(ctx context.Context, db *sqlx.DB, driverId int64) (*models.Driver, error) {
	var result models.Driver

	query, args, _ := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Where(sq.Eq{"id": driverId}).
		ToSql()

	err := db.GetContext(ctx, &result, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"DriverRepository", "GetDriverByID"}).
			Str("query", query).
			Msg(err.Error())

		return nil, err
	}

	return &result, nil

}
