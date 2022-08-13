package origins

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type OriginRepositoryInterface interface {
	GetTableName() string
	GetOrigins(ctx context.Context, db *sqlx.DB, filter *models.OriginFilter) ([]models.Origin, error)
	// GetOriginByID(ctx context.Context, db *sqlx.DB, id int64) (*models.Origin, error)
	// GetOriginByName(ctx context.Context, db *sqlx.DB, name string) (*models.Origin, error)
	// Create(ctx context.Context, db *sqlx.DB, name string) (*models.Origin, error)
}

type originRepository struct {
	dbName string
}

func NewRepository(dbName string) OriginRepositoryInterface {
	return &originRepository{
		dbName: dbName,
	}
}

func (repo *originRepository) selectFields() []string {
	return []string{
		"id",
		"name",
	}
}

func (repo *originRepository) GetTableName() string {
	return "origins"
}
