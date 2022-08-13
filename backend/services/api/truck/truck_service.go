package truck

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type truckService struct {
	db   *sqlx.DB
	repo TruckRepositoryInterface
}

func NewService(db *sqlx.DB, repo TruckRepositoryInterface) TruckServiceInterface {
	return &truckService{
		db:   db,
		repo: repo,
	}
}

func (svc *truckService) GetTrucks(ctx context.Context, filter *models.TruckFilter) ([]models.Truck, error) {
	trucks, err := svc.repo.GetTrucks(ctx, svc.db, filter)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTrucks"}).
			Msg(err.Error())
		return nil, err
	}

	return trucks, nil
}

func (svc *truckService) GetTruckByID(ctx context.Context, truckID int64) (*models.Truck, error) {
	truck, err := svc.repo.GetTruckByID(ctx, svc.db, truckID)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTruck"}).
			Msg(err.Error())
		return nil, err
	}

	return truck, nil
}
