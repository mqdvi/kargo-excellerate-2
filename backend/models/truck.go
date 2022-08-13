package models

import (
	"github.com/gin-gonic/gin"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
)

type Truck struct {
	ID             int64   `db:"id" json:"id"`
	TruckType      string  `db:"truck_type" json:"truckType"`
	LicenseNumber  string  `db:"license_number" json:"licenseNumber"`
	PlateType      string  `db:"plate_type" json:"plateType"`
	ProductionYear int64   `db:"production_year" json:"productionYear"`
	Status         bool    `db:"status" json:"status"`
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
