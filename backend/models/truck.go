package models

import (
	"github.com/gin-gonic/gin"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
)

const (
	TRUCK_STATUS_ACTIVE   = "ACTIVE"
	TRUCK_STATUS_INACTIVE = "INACTIVE"
)

type Truck struct {
	ID             int64   `db:"id" json:"id"`
	TruckType      string  `db:"truck_type" json:"truckType"`
	LicenseNumber  string  `db:"license_number" json:"licenseNumber"`
	PlateType      string  `db:"plate_type" json:"plateType"`
	ProductionYear int64   `db:"production_year" json:"productionYear"`
	IsActive       bool    `db:"status" json:"-"`
	Status         string  `db:"-" json:"status"`
	STNK           *string `db:"stnk" json:"stnk"`
	KIR            *string `db:"kir" json:"kir"`
}

type TruckFilter struct {
	Search      string
	Sort        string
	TruckTypeID int64
}

func (m *TruckFilter) FromContext(c *gin.Context) *TruckFilter {
	m.Search = c.Query("search")
	m.Sort = c.Query("sort")
	m.TruckTypeID = int64(helper.StringToInt(c.Query("truckTypeId")))

	return m
}

type CreateTruckPayload struct {
	LicenseNumber  string  `json:"licenseNumber" validate:"required"`
	PlateType      string  `json:"plateType" validate:"required,oneof=Yellow Black"`
	TruckTypeID    int64   `json:"truckTypeId" validate:"required"`
	ProductionYear int64   `json:"productionYear" validate:"required"`
	STNK           *string `json:"stnk"`
	STNKUrl        string  `json:"-"`
	KIR            *string `json:"kir"`
	KirUrl         string  `json:"-"`
}
