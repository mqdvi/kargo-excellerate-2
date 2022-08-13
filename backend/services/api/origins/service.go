package origins

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type OriginServiceInterface interface {
	GetOrigins(ctx context.Context, filter *models.OriginFilter) (*[]models.Origin, error)
}

type originService struct {
	db   *sqlx.DB
	repo OriginRepositoryInterface
}

func NewService(db *sqlx.DB, repo OriginRepositoryInterface) OriginServiceInterface {
	return &originService{
		db:   db,
		repo: repo,
	}
}
