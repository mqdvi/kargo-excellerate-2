package truck_type

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type TruckTypeRepositoryInterface interface {
	GetTableName() string
	GetTruckTypes(ctx context.Context, db *sqlx.DB) ([]models.TruckType, error)
}

type TruckTypeServiceInterface interface {
	GetTruckTypes(ctx context.Context) ([]models.TruckType, error)
}
