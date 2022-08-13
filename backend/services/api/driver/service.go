package driver

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type DriverServiceInterface interface {
	GetDriver(ctx context.Context, payload *models.GetDriverFilter) (*models.ResponseGetDrivers, error)
	Create(ctx context.Context, payload *models.RequestCreateDriver) (*models.ResponseCreateDriver, error)
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