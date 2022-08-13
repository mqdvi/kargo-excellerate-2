package truck

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type TruckRepositoryInterface interface {
	GetTableName() string
	GetTrucks(ctx context.Context, db *sqlx.DB, filter *models.TruckFilter) ([]models.Truck, error)
	GetTruckByID(ctx context.Context, db *sqlx.DB, truckID int64) (*models.Truck, error)
}

type TruckServiceInterface interface {
	GetTrucks(ctx context.Context, filter *models.TruckFilter) ([]models.Truck, error)
	GetTruckByID(ctx context.Context, truckID int64) (*models.Truck, error)
}
