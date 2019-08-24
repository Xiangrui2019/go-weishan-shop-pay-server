package services

import (
	"go-weishan-shop-pay-server/models"

	"github.com/gin-gonic/gin"
)

type PublishOrderService struct {
}

func (service *PublishOrderService) Publish(context *gin.Context) string {
	err := models.PublishOrder(context.Param("id"))

	if err != nil {
		return ""
	}

	return "<h1 style='font-size: 100px'>发货成功!</h1>"
}
