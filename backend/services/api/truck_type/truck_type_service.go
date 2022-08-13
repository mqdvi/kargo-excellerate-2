package truck_type

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type truckTypeService struct {
	db   *sqlx.DB
	repo TruckTypeRepositoryInterface
}

func NewService(db *sqlx.DB, repo TruckTypeRepositoryInterface) TruckTypeServiceInterface {
	return &truckTypeService{
		db:   db,
		repo: repo,
	}
}

func (svc *truckTypeService) GetTruckTypes(ctx context.Context) ([]models.TruckType, error) {
	truckTypes, err := svc.repo.GetTruckTypes(ctx, svc.db)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckTypeRepository", "GetTruckTypes"}).
			Msg(err.Error())
		return nil, err
	}

	return truckTypes, nil
}
