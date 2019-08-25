package services

import (
	"encoding/json"
	"go-weishan-shop-pay-server/cache"
	"go-weishan-shop-pay-server/global"
	"go-weishan-shop-pay-server/modules"
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/utils"
	"net/http"
	"os"
)

type CreatePayService struct {
	Goodname    string  `form:"good_name" json:"good_name" binding:"required"`
	GoodId      string  `form:"good_id" json:"good_id" binding:"required"`
	Realname    string  `form:"real_name" json:"real_name" binding:"required,min=2,max=20"`
	Address     string  `form:"address" json:"address" binding:"required,min=5,max=100"`
	Phonenumber string  `form:"phone_number" json:"phone_number" binding:"required,min=11,max=21"`
	ExtInfo     string  `form:"extinfo" json:"extinfo"`
	BuyCount    int     `form:"buy_count" json:"buy_count" binding:"required"`
	BuyPrice    float64 `form:"buy_price" json:"buy_price" binding:"required"`
	Price       float64 `form:"price" json:"price" binding:"required"`
}

func (service *CreatePayService) buildOrderCache() (string, error) {
	value, err := json.Marshal(&global.OrderCache{
		Goodname:    service.Goodname,
		GoodId:      service.GoodId,
		Realname:    service.Realname,
		Address:     service.Address,
		Phonenumber: service.Phonenumber,
		ExtInfo:     service.ExtInfo,
		BuyCount:    service.BuyCount,
		BuyPrice:    service.BuyPrice,
	})

	if err != nil {
		return "", err
	}

	return string(value), err
}

func (service *CreatePayService) buildCashier(token string) string {
	cashier := modules.PayJSModule.GetCashier()
	url, err := cashier.GetRequestUrl(
		utils.Yuan2fen(utils.CalcGoodPrice(service.Price, service.BuyCount)),
		"微山掌上拍付款",
		utils.RandomString(32),
		token,
		os.Getenv("APP_CALLBACK_PAGE"),
		1,
		1,
	)

	if err != nil {
		panic(err)
	}

	return url
}

func (service *CreatePayService) Create() (string, *serializer.Response) {
	token := utils.RandomString(64)
	value, err := service.buildOrderCache()

	if err != nil {
		return "", &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "转换数据出错.",
			Error:   err.Error(),
		}
	}

	_, err = cache.CacheClient.Set(token, value, 0).Result()

	if err != nil {
		return "", &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "Redis 写入出错.",
			Error:   err.Error(),
		}
	}

	url := service.buildCashier(token)

	return url, nil
}
