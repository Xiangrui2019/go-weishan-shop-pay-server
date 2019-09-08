package api

import (
	"go-weishan-shop-pay-server/modules"
	"go-weishan-shop-pay-server/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	err := modules.HealthChecksModule.Check()

	if len(err) != 0 {
		context.JSON(http.StatusInternalServerError, &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "HealthCheck出错",
			Data:    err,
		})
		return
	}

	context.JSON(http.StatusOK, &serializer.Response{
		Code:    http.StatusOK,
		Message: "Pong",
	})
}
