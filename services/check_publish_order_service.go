package services

import (
	"go-weishan-shop-pay-server/models"

	"github.com/gin-gonic/gin"
)

type CheckPublishOrderService struct {
}

func (service *CheckPublishOrderService) CheckPublish(context *gin.Context) string {
	order, err := models.GetOrderById(context.Param("id"))

	if err != nil {
		return ""
	}

	if order.Status == false {
		return "<h1 style='color: red; font-size: 100px'>未发货!</h1>"
	} else {
		return "<h1 style='color: green; font-size: 100px'>已发货!</h1>"
	}
}
