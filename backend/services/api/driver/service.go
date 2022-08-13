package driver

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type DriverServiceInterface interface {
	GetDriver(ctx context.Context, payload *models.GetDriverFilter) (*models.ResponseGetDrivers, error)
	GetDriverByID(ctx context.Context, driverID int64) (*models.Driver, error)
	Create(ctx context.Context, payload *models.RequestCreateDriver) (*models.ResponseCreateDriver, error)
	Update(ctx context.Context, payload *models.RequestCreateDriver, driverID int64) (*models.ResponseCreateDriver, error)
	Deactivate(ctx context.Context, driverID int64) (*models.Driver, error)
}

type driverService struct {
	db   *sqlx.DB
	repo DriverRepositoryInterface
}

func NewService(db *sqlx.DB, repo DriverRepositoryInterface) DriverServiceInterface {
	return &driverService{
		db:   db,
		repo: repo,
	}
}
