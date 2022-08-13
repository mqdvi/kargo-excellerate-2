package driver

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (repo *driverRepository) GetDriver(ctx context.Context, db *sqlx.DB, payload *models.GetDriverFilter) ([]models.Driver, error) {
	var result []models.Driver
	query, args, _ := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).ToSql()

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