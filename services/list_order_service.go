package services

import (
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/serializer"
	"net/http"
)

type ListOrderService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

func (service *ListOrderService) List() *serializer.Response {
	orders := []models.Order{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := models.DB.Model(models.Order{}).Count(&total).Error; err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "数据库连接错误",
			Error:   err.Error(),
		}
	}

	if err := models.DB.Limit(service.Limit).Offset(service.Start).Find(&orders).Error; err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "数据库连接错误",
			Error:   err.Error(),
		}
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "获取订单列表成功.",
		Data: serializer.BuildDataList(
			serializer.BuildOrders(orders),
			uint(total)),
	}
}
