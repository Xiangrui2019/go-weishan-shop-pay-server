package services

import (
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublishOrderService struct {
}

func (service *PublishOrderService) Publish(context *gin.Context) *serializer.Response {
	err := models.PublishOrder(context.Param("id"))

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "写入数据库出错",
			Error:   err.Error(),
		}
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "发布商品成功.",
	}
}
