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
	GetTruckByLicenseNumber(ctx context.Context, db *sqlx.DB, licenseNumber string, truckID *int64) (*models.Truck, error)
	Create(ctx context.Context, db *sqlx.DB, payload *models.CreateTruckPayload) error
	Update(ctx context.Context, db *sqlx.DB, payload *models.CreateTruckPayload, truckID int64) error
	DeactivateTruck(ctx context.Context, db *sqlx.DB, truckID int64) error
}

type TruckServiceInterface interface {
	GetTrucks(ctx context.Context, filter *models.TruckFilter) ([]models.Truck, error)
	GetTruckByID(ctx context.Context, truckID int64) (*models.Truck, error)
	Create(ctx context.Context, payload *models.CreateTruckPayload) error
	Update(ctx context.Context, payload *models.CreateTruckPayload, truckID int64) error
	DeactivateTruck(ctx context.Context, truckID int64) error
}
