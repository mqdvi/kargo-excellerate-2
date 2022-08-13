package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (ctrl *driverController) HandlerDeactivateDriver(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))

	result, err := ctrl.driverSvc.DeactivateDriver(c, int64(id))
	if err != nil {
		errorResponse := helper.NewErrorResponse("driver-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}
	
	response := helper.DataResponse{
		Data: result,
	}

	c.JSON(http.StatusOK, response)
}
