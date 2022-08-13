package shipment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
)

type shipmentController struct {
	shipmentSvc ShipmentServiceInterface
	validator   *validator.Validate
}

func NewController(shipmentSvc ShipmentServiceInterface, validator *validator.Validate) *shipmentController {
	return &shipmentController{
		shipmentSvc: shipmentSvc,
		validator:   validator,
	}
}

func (ctrl *shipmentController) HandlerGetShipments(c *gin.Context) {
	filter := models.ShipmentFilter{}
	filter.FromContext(c)

	result, err := ctrl.shipmentSvc.GetShipments(c, &filter)
	if err != nil {
		errorResponse := helper.NewErrorResponse("shipments-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.ItemsReponse{}
	response.Data.Items = result

	c.JSON(http.StatusOK, response)
}

func (ctrl *shipmentController) HandlerCreateShipment(c *gin.Context) {
	payload := new(models.CreateShipmentPayload)
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorResponse := helper.NewErrorResponse("shipments-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err := ctrl.validator.Struct(payload); err != nil {
		errorResponse := helper.NewErrorResponse("shipments-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := ctrl.shipmentSvc.CreateShipment(c, payload)
	if err != nil {
		errorResponse := helper.NewErrorResponse("shipments-500", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: "OK",
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *shipmentController) HandlerAllocateShipment(c *gin.Context) {
	payload := new(models.AllocateShipmentPayload)
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorResponse := helper.NewErrorResponse("shipments-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err := ctrl.validator.Struct(payload); err != nil {
		errorResponse := helper.NewErrorResponse("shipments-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	id := helper.StringToInt(c.Param("id"))
	err := ctrl.shipmentSvc.AllocateShipment(c, payload, int64(id))
	if err != nil {
		errorResponse := helper.NewErrorResponse("shipments-500", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: "OK",
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *shipmentController) HandlerUpdateShipmentStatus(c *gin.Context) {
	payload := new(models.UpdateShipmentStatusPayload)
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorResponse := helper.NewErrorResponse("shipments-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err := ctrl.validator.Struct(payload); err != nil {
		errorResponse := helper.NewErrorResponse("shipments-400", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	id := helper.StringToInt(c.Param("id"))
	err := ctrl.shipmentSvc.UpdateStatus(c, payload, int64(id))
	if err != nil {
		errorResponse := helper.NewErrorResponse("shipments-500", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.DataResponse{
		Data: "OK",
	}

	c.JSON(http.StatusCreated, response)
}
