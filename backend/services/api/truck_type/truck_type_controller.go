package truck_type

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
)

type truckTypeController struct {
	truckTypeSvc TruckTypeServiceInterface
}

func NewController(truckTypeSvc TruckTypeServiceInterface) *truckTypeController {
	return &truckTypeController{truckTypeSvc: truckTypeSvc}
}

func (ctrl *truckTypeController) HandlerGetTruckTypes(c *gin.Context) {
	result, err := ctrl.truckTypeSvc.GetTruckTypes(c)
	if err != nil {
		errorResponse := helper.NewErrorResponse("truckTypes-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.ItemsReponse{}
	response.Data.Items = result

	c.JSON(http.StatusOK, response)
}
