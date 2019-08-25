package services

import (
	"fmt"
	"go-weishan-shop-pay-server/cache"
	"go-weishan-shop-pay-server/global"
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/modules"
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/templates"
	"go-weishan-shop-pay-server/utils"
	"net/http"
	"os"
	"strconv"

	"encoding/json"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/qingwg/payjs/notify"
)

type ConfirmPayService struct {
}

func (service *ConfirmPayService) getOrderCache(msg notify.Message) (*global.OrderCache, error) {
	data := global.OrderCache{}

	value, err := cache.CacheClient.Get(msg.Attach).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (service *ConfirmPayService) buildOrder(data *global.OrderCache) *models.Order {
	return &models.Order{
		Goodname:    data.Goodname,
		GoodId:      data.GoodId,
		Realname:    data.Realname,
		Address:     data.Address,
		Phonenumber: data.Phonenumber,
		ExtInfo:     data.ExtInfo,
		BuyCount:    data.BuyCount,
		BuyPrice:    data.BuyPrice,
		Status:      false,
	}
}

func (service *ConfirmPayService) buildFee(data *global.OrderCache, to float64, fee float64) *models.Fee {
	return &models.Fee{
		TotalValue: data.BuyPrice,
		ToValue:    to,
		FeeValue:   fee,
	}
}

func (service *ConfirmPayService) buildWeixinContent(data *global.OrderCache, order *models.Order) string {
	return fmt.Sprintf(templates.WeixinNotifySenderTemplate,
		data.Goodname,
		data.GoodId,
		data.Realname,
		data.Address,
		data.Phonenumber,
		data.ExtInfo,
		data.BuyCount,
		data.BuyPrice,
		os.Getenv("CHECK_CALLBACK"),
		strconv.Itoa(int(order.ID)),
		os.Getenv("APP_MAP"),
		url.QueryEscape(data.Address),
		os.Getenv("FINISH_CALLBACK"),
		strconv.Itoa(int(order.ID)))
}

func (service *ConfirmPayService) messageHandler(context *gin.Context, msg notify.Message) {
	data, err := service.getOrderCache(msg)

	if err != nil {
		panic(err)
	}

	order, err := models.CreateOrder(service.buildOrder(data))
	if err != nil {
		panic(err)
	}

	feerate, err := strconv.ParseFloat(os.Getenv("FEE_RATE"), 64)
	to, fee := utils.CalcFee(data.BuyPrice, feerate)

	err = models.CreateFee(service.buildFee(data, to, fee))
	if err != nil {
		panic(err)
	}

	utils.SendWeixinNotify(
		"收款消息提醒",
		service.buildWeixinContent(data, order),
		"nourl",
	)

	_, err = cache.CacheClient.Del(msg.Attach).Result()
	if err != nil {
		panic(err)
	}
}

func (service *ConfirmPayService) Finish(context *gin.Context) *serializer.Response {
	payNotify := modules.PayJSModule.GetNotify(context.Request, context.Writer)

	payNotify.SetMessageHandler(func(msg notify.Message) {
		service.messageHandler(context, msg)
	})

	err := payNotify.Serve()

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "出错啦!",
			Error:   err.Error(),
		}
	}

	return nil
}
