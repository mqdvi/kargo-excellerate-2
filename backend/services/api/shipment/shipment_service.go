package shipment

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/api/driver"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/api/truck"
)

type shipmentService struct {
	db         *sqlx.DB
	repo       ShipmentRepositoryInterface
	truckRepo  truck.TruckRepositoryInterface
	driverRepo driver.DriverRepositoryInterface
}

func NewService(db *sqlx.DB, repo ShipmentRepositoryInterface, truckRepo truck.TruckRepositoryInterface, driverRepo driver.DriverRepositoryInterface) ShipmentServiceInterface {
	return &shipmentService{
		db:         db,
		repo:       repo,
		truckRepo:  truckRepo,
		driverRepo: driverRepo,
	}
}

func (svc *shipmentService) GetShipments(ctx context.Context, filter *models.ShipmentFilter) (*[]models.Shipment, error) {
	shipments, err := svc.repo.GetShipments(ctx, svc.db, filter)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"ShipmentRepository", "GetShipments"}).
			Msg(err.Error())

		return nil, err
	}

	return &shipments, nil
}

func (svc *shipmentService) CreateShipment(ctx context.Context, payload *models.CreateShipmentPayload) error {
	shipmentNumber := "DO-" + fmt.Sprint(rand.Intn(10000))
	payload.ShipmentNumber = shipmentNumber

	err := svc.repo.CreateShipment(ctx, svc.db, payload)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"ShipmentRepository", "CreateShipment"}).
			Msg(err.Error())

		return err
	}

	return nil
}

func (svc *shipmentService) AllocateShipment(ctx context.Context, payload *models.AllocateShipmentPayload, shipmentId int64) error {
	truck, _ := svc.truckRepo.GetTruckByID(ctx, svc.db, payload.TruckID)
	if truck == nil {
		return errors.New("Truck not found")
	}

	driver, _ := svc.driverRepo.GetDriverByID(ctx, svc.db, payload.DriverID)
	if driver == nil {
		return errors.New("Driver not found")
	}

	err := svc.repo.AllocateShipment(ctx, svc.db, payload, shipmentId)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"ShipmentRepository", "AllocateShipment"}).
			Msg(err.Error())

		return err
	}

	return nil
}

func (svc *shipmentService) UpdateStatus(ctx context.Context, payload *models.UpdateShipmentStatusPayload, shipmentId int64) error {
	err := svc.repo.UpdateStatus(ctx, svc.db, payload, shipmentId)
	if err != nil {
		helper.Logger.Error().
			Strs("tags", []string{"ShipmentRepository", "UpdateStatus"}).
			Msg(err.Error())
		return err
	}

	return nil
}
