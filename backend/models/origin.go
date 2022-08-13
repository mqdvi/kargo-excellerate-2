package models

import (
	"github.com/gin-gonic/gin"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
)

type Origin struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type OriginFilter struct {
	Search      string
	ID        	int64
}

func (m *OriginFilter) FromContext(c *gin.Context) *OriginFilter {
	m.Search = c.Query("search")
	m.ID = int64(helper.StringToInt(c.Query("Id")))

	return m
}