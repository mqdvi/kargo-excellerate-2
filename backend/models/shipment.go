package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SHIPMENT_STATUS_ONGOING_TO_ORIGIN      = "Ongoing to Origin"
	SHIPMENT_STATUS_AT_ORIGIN              = "At Origin"
	SHIPMENT_STATUS_ONGOING_TO_DESTINATION = "Ongoint to Destination"
	SHIPMENT_STATUS_AT_DESTINATION         = "At Destination"
	SHIPMENT_STATUS_COMPLETED              = "COMPLETED"
)

type Shipment struct {
	ID                  int64   `db:"id" json:"id"`
	ShipmentNumber      string  `db:"shipment_number" json:"shipmentNumber"`
	LicenseNumber       *string `db:"license_number" json:"licenseNumber"`
	DriverName          *string `db:"driver_name" json:"driverName"`
	OriginDistrict      string  `db:"origin_district" json:"originDistrict"`
	DestinationDistrict string  `db:"destination_district" json:"destinationDistrict"`
}

type ShipmentFilter struct {
	Search string
	Sort   string
}

func (m *ShipmentFilter) FromContext(c *gin.Context) *ShipmentFilter {
	m.Search = c.Query("search")
	m.Sort = c.Query("sort")

	return m
}

type CreateShipmentPayload struct {
	OriginID       int64     `json:"originId" validate:"required"`
	DestinationID  int64     `json:"destinationId" validate:"required"`
	LoadingDate    time.Time `json:"loadingDate" validate:"required"`
	ShipmentNumber string    `json:"-"`
}

type AllocateShipmentPayload struct {
	TruckID  int64 `json:"truckId" validate:"required"`
	DriverID int64 `json:"driverId" validate:"required"`
}

type UpdateShipmentStatusPayload struct {
	Status string `json:"status" validate:"required,oneof='Ongoing to Origin' 'At Origin' 'Ongoing to Destination' 'At Destination' 'Completed'"`
}
