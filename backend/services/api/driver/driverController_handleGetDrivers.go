package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	"net/http"
)

func (ctrl *driverController) HandlerGetDrivers(c *gin.Context) {

	filter := models.GetDriverFilter{}
	filter.FromContext(c)

	result, err := ctrl.driverSvc.GetDriver(c, &filter)
	if err != nil {
		errorResponse := helper.NewErrorResponse("driver-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, result)
}
