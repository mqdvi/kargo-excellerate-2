package driver

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type DriverRepositoryInterface interface {
	CreateDriver(ctx context.Context, db *sqlx.DB, payload *models.RequestCreateDriver) (*models.ResponseCreateDriver, error)
	GetDriver(ctx context.Context, db *sqlx.DB, payload *models.GetDriverFilter) ([]models.Driver, error)
	GetDriverByID(ctx context.Context, db *sqlx.DB, driverID int64) (*models.Driver, error)
	// UpdateDriver(ctx context.Context, db *sqlx.DB, payload *models.GetDriverFilter) ([]models.Driver, error)
}
type driverRepository struct {
	dbName string
}

func (repo *driverRepository) GetTableName() string {
	return "drivers"
}

func (repo *driverRepository) selectFields() []string {
	return []string{
		"drivers.id",
		"drivers.name",
		"drivers.phone_number",
		"drivers.created_at",
		"drivers.status",
	}
}

func NewRepository(dbName string) DriverRepositoryInterface {
	return &driverRepository{
		dbName: dbName,
	}
}
