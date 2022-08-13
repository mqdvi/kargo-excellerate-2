package origins

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (repo *originRepository) GetOrigins(ctx context.Context, db *sqlx.DB, filter *models.OriginFilter) ([]models.Origin, error) {
	var result []models.Origin

	builder := sq.Select(repo.selectFields()...).
		From(repo.GetTableName())

	if filter != nil {
		if filter.Search != "" {
			builder = builder.Where(sq.Or{
				sq.Like{"name": "%" + filter.Search + "%"},
			})
		}
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
