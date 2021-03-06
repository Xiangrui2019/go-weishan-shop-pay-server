package api

import (
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/services"
	"go-weishan-shop-pay-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePay(context *gin.Context) {
	service := services.CreatePayService{}

	if err := context.ShouldBind(&service); err != nil {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	} else {
		if url, err := service.Create(); err == nil {
			context.JSON(http.StatusOK, &serializer.Response{
				Code:    http.StatusOK,
				Message: "订单创建成功.",
				Data:    url,
			})
		} else {
			context.JSON(http.StatusInternalServerError, err)
		}
	}
}

func ConfirmPay(context *gin.Context) {
	service := services.ConfirmPayService{}

	if err := service.Finish(context); err == nil {
		context.JSON(http.StatusOK, &serializer.Response{
			Code:    http.StatusOK,
			Message: "订单确认成功.",
		})
	} else {
		context.JSON(http.StatusInternalServerError, err)
	}
}
