package origins

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

func (repo *originRepository) GetOriginByName(ctx context.Context, db *sqlx.DB, name string) (*models.Origin, error) {
	query, args, _ := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		Where(sq.Eq{"name": name}).
		Limit(1).
		ToSql()

	var result models.Origin
	err := db.GetContext(ctx, &result, query, args...)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
