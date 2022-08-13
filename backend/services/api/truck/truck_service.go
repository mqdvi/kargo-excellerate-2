package truck

import (
	"context"
	"database/sql"
	"errors"

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

	for i, truck := range trucks {
		if truck.IsActive {
			trucks[i].Status = models.TRUCK_STATUS_ACTIVE
		} else {
			trucks[i].Status = models.TRUCK_STATUS_INACTIVE
		}
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

	if truck.IsActive {
		truck.Status = models.TRUCK_STATUS_ACTIVE
	} else {
		truck.Status = models.TRUCK_STATUS_INACTIVE
	}

	return truck, nil
}

func (svc *truckService) Create(ctx context.Context, payload *models.CreateTruckPayload) error {
	truck, err := svc.repo.GetTruckByLicenseNumber(ctx, svc.db, payload.LicenseNumber, nil)
	if err != nil && err != sql.ErrNoRows {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTruckByLicenseNumber"}).
			Msg(err.Error())
		return err
	}

	if truck != nil {
		return errors.New("Invalid license number, truck already exists.")
	}

	err = svc.repo.Create(ctx, svc.db, payload)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "Create"}).
			Msg(err.Error())
		return err
	}

	return nil
}

func (svc *truckService) Update(ctx context.Context, payload *models.CreateTruckPayload, truckID int64) error {
	truck, err := svc.repo.GetTruckByLicenseNumber(ctx, svc.db, payload.LicenseNumber, &truckID)
	if err != nil && err != sql.ErrNoRows {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "GetTruckByLicenseNumber"}).
			Msg(err.Error())
		return err
	}

	if truck != nil {
		return errors.New("Invalid license number, truck already exists.")
	}

	err = svc.repo.Update(ctx, svc.db, payload, truckID)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "Update"}).
			Msg(err.Error())
		return err
	}

	return nil
}
	
func (svc *truckService) DeactivateTruck(ctx context.Context, truckID int64) error {
	err := svc.repo.DeactivateTruck(ctx, svc.db, truckID)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"TruckRepository", "DeactivateTruck"}).
			Msg(err.Error())
		return err
	}

	return nil
}
