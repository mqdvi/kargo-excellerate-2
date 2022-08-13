package shipment

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type ShipmentRepositoryInterface interface {
	GetTableName() string
	GetShipments(ctx context.Context, db *sqlx.DB, filter *models.ShipmentFilter) ([]models.Shipment, error)
	CreateShipment(ctx context.Context, db *sqlx.DB, payload *models.CreateShipmentPayload) error
	AllocateShipment(ctx context.Context, db *sqlx.DB, payload *models.AllocateShipmentPayload, shipmentId int64) error
	UpdateStatus(ctx context.Context, db *sqlx.DB, payload *models.UpdateShipmentStatusPayload, shipmentId int64) error	
}

type ShipmentServiceInterface interface{
	GetShipments(ctx context.Context, filter *models.ShipmentFilter) (*[]models.Shipment, error)
	CreateShipment(ctx context.Context, payload *models.CreateShipmentPayload) error
	AllocateShipment(ctx context.Context, payload *models.AllocateShipmentPayload, shipmentId int64) error
	UpdateStatus(ctx context.Context, payload *models.UpdateShipmentStatusPayload, shipmentId int64) error	
}
