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

func NewController(truckSvc TruckServiceInterface, validator *validator.Validate) *truckController {
	return &truckController{
		truckSvc:  truckSvc,
		validator: validator,
	}
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

func (ctrl *truckController) HandlerCreateTruck(c *gin.Context) {
	payload := new(models.CreateTruckPayload)
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorResponse := helper.NewErrorResponse("trucks-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err := ctrl.validator.Struct(payload); err != nil {
		errorResponse := helper.NewErrorResponse("trucks-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := ctrl.truckSvc.Create(c, payload)
	if err != nil {
		errorResponse := helper.NewErrorResponse("trucks-500", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: "OK",
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *truckController) HandlerUpdateTruck(c *gin.Context) {
	payload := new(models.CreateTruckPayload)
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorResponse := helper.NewErrorResponse("trucks-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err := ctrl.validator.Struct(payload); err != nil {
		errorResponse := helper.NewErrorResponse("trucks-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	id := helper.StringToInt(c.Param("id"))
	err := ctrl.truckSvc.Update(c, payload, int64(id))
	if err != nil {
		errorResponse := helper.NewErrorResponse("trucks-500", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: "OK",
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *truckController) HandlerDeactivateTruck(c *gin.Context) {
	id := helper.StringToInt(c.Param("id"))
	err := ctrl.truckSvc.DeactivateTruck(c, int64(id))
	if err != nil {
		errorResponse := helper.NewErrorResponse("trucks-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: "OK",
	}

	c.JSON(http.StatusOK, response)
}
