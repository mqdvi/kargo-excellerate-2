package origins

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mqdvi/kargo-excellerate-2/backend/helper"
	"github.com/mqdvi/kargo-excellerate-2/backend/models"
	// "github.com/mqdvi/kargo-excellerate-2/backend/services/config"
)

func (ctrl *originController) HandlerGetOrigins(c *gin.Context) {
	filter := models.OriginFilter{}
	filter.FromContext(c)

	result, err := ctrl.originSvc.GetOrigins(c, &filter)
	if err != nil {
		errorResponse := helper.NewErrorResponse("origins-500", err.Error())

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.ItemsReponse{}
	response.Data.Items = result

	c.JSON(http.StatusOK, response)
}
