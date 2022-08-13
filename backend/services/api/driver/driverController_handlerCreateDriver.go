package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	"net/http"
)

func (ctrl *driverController) HandlerCreateArticle(c *gin.Context) {
	payload := new(models.RequestCreateDriver)
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorResponse := helper.NewErrorResponse("driver-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err := ctrl.validator.Struct(payload); err != nil {
		errorResponse := helper.NewErrorResponse("driver-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	result, err := ctrl.driverSvc.Create(c, payload)
	if err != nil {
		errorResponse := helper.NewErrorResponse("driver-500", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.JsonResponse{
		Data: result,
	}

	c.JSON(http.StatusCreated, response)
}
