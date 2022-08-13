package origins

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type OriginServiceInterface interface {
	GetOriginsByName(ctx context.Context) (*Data, error)
	// GetOriginsPagination(ctx context.Context, filter *models.GetOriginsFilter, meta *helper.Meta) (*helper.PaginationResponse, error)
	// Create(ctx context.Context, payload *models.CreateOriginPayload) (*models.Origin, error)
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
