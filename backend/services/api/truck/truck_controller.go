package truck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type truckController struct {
	truckSvc  TruckServiceInterface
	validator *validator.Validate
}

func NewController(truckSvc TruckServiceInterface) *truckController {
	return &truckController{truckSvc: truckSvc}
}

func (ctrl *truckController) HandlerGetTrucks(c *gin.Context) {
	filter := models.TruckFilter{}
	filter.FromContext(c)

	result, err := ctrl.truckSvc.GetTrucks(c, &filter)
	if err != nil {
		errorResponse := helper.NewErrorResponse("trucks-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.ItemsReponse{}
	response.Data.Items = result

	c.JSON(http.StatusOK, response)
}

func (ctrl *truckController) HandlerGetTruckByID(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))

	result, err := ctrl.truckSvc.GetTruckByID(c, int64(id))
	if err != nil {
		errorResponse := helper.NewErrorResponse("trucks-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: result,
	}

	c.JSON(http.StatusOK, response)
}
