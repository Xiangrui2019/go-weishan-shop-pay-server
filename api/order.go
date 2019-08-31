package api

import (
	"go-weishan-shop-pay-server/services"
	"go-weishan-shop-pay-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrder(context *gin.Context) {
	service := services.ListOrderService{}

	if err := context.ShouldBind(&service); err == nil {
		result := service.List()

		context.JSON(result.Code, result)
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func ListNonPubOrder(context *gin.Context) {
	service := services.ListNonPubOrderService{}

	if err := context.ShouldBind(&service); err == nil {
		result := service.List()

		context.JSON(result.Code, result)
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}

func PublishOrder(context *gin.Context) {
	service := services.PublishOrderService{}

	res := service.Publish(context)
	context.JSON(res.Code, res)
}
