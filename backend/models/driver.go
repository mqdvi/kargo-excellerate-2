package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Driver struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	PhoneNumber string    `db:"phone_name" json:"phoneNumber"`
	Status      string    `db:"status" json:"status,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt,omitempty"`
}
type ResponseGetDrivers struct {
	Data DriverPayloadData `json:"data"`
}
type DriverPayloadData struct {
	Items []Driver `json:"items"`
}

type RequestCreateDriver struct {
	Name          string `json:"name"`
	PhoneNumber   string `json:"phoneNumber"`
	IDCard        string `json:"idCard"`
	DriverLicense string `json:"driverLicense"`
}
type ResponseCreateDriver struct {
	Data string `json:"data"`
}

type GetDriverFilter struct {
	Search string
	Sort   string
}

func (m *GetDriverFilter) FromContext(c *gin.Context) *GetDriverFilter {
	m.Search = c.Query("search")
	m.Sort = c.Query("sort")

	return m
}
