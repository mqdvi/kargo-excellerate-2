package truck_type

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type truckTypeRepository struct {
	dbName string
}

func NewRepository(dbName string) TruckTypeRepositoryInterface {
	return &truckTypeRepository{
		dbName: dbName,
	}
}

func (repo *truckTypeRepository) selectFields() []string {
	return []string{
		"id",
		"name",
	}
}

func (repo *truckTypeRepository) GetTableName() string {
	return "truck_types"
}

func (repo *truckTypeRepository) GetTruckTypes(ctx context.Context, db *sqlx.DB) ([]models.TruckType, error) {
	var result []models.TruckType

	query, args, _ := sq.Select(repo.selectFields()...).
		From(repo.GetTableName()).
		ToSql()

	err := db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckTypeRepository", "GetTruckTypes"}).
			Str("query", query).
			Msg(err.Error())

		return result, err
	}

	return result, nil
}
