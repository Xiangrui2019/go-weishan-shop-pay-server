package api

import (
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrder(context *gin.Context) {
	service := services.ListOrderService{}

	if orders, err := service.List(); err == nil {
		context.JSON(http.StatusOK, &serializer.Response{
			Code:    http.StatusOK,
			Message: "成功获取订单列表.",
			Data:    orders,
		})
	} else {
		context.JSON(http.StatusInternalServerError, err)
	}
}

func PublishOrder(context *gin.Context) {
	service := services.PublishOrderService{}

	if err := service.Publish(context); err != "" {
		context.Writer.Header().Set("Content-Type", "text/html;charset=utf-8")
		context.Writer.Write([]byte(err))
	} else {
		context.String(http.StatusInternalServerError, "服务器出错啦!", nil)
	}
}

func CheckPublishOrder(context *gin.Context) {
	service := services.CheckPublishOrderService{}

	if err := service.CheckPublish(context); err != "" {
		context.Writer.Header().Set("Content-Type", "text/html;charset=utf-8")
		context.Writer.Write([]byte(err))
	} else {
		context.String(http.StatusInternalServerError, "服务器出错!")
	}
}
