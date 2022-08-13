package origins

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (repo *originRepository) GetOriginByID(ctx context.Context, db *sqlx.DB, id int64) (*models.Origin, error) {
	query, args, _ := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Where(sq.Eq{"id": id}).
		Limit(1).
		ToSql()

	var result models.Origin
	err := db.GetContext(ctx, &result, query, args...)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
