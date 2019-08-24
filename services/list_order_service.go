package services

import (
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/serializer"
	"net/http"
)

type ListOrderService struct {
}

func (service *ListOrderService) List() ([]models.Order, *serializer.Response) {
	orders, err := models.ListOrder()

	if err != nil {
		return nil, &serializer.Response{
			Code:    http.StatusOK,
			Message: "获取错误.",
			Error:   err.Error(),
		}
	}

	return orders, nil
}
