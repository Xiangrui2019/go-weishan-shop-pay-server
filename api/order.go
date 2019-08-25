package api

import (
	"go-weishan-shop-pay-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
