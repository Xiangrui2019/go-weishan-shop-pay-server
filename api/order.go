package api

import (
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/services"
	"go-weishan-shop-pay-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(context *gin.Context) {
	service := services.CreateOrderService{}

	if err := context.ShouldBind(&service); err != nil {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	} else {
		if payurl, err := service.Create(); err == nil {
			context.JSON(http.StatusOK, &serializer.Response{
				Code:    http.StatusOK,
				Message: "订单创建成功.",
				Data:    payurl,
			})
		} else {
			context.JSON(http.StatusInternalServerError, err)
		}
	}
}

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

func FinishOrder(context *gin.Context) {
	service := services.FinishOrderService{}

	if err := service.Finish(context); err == nil {
		context.JSON(http.StatusOK, &serializer.Response{
			Code:    http.StatusOK,
			Message: "订单确认成功.",
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
