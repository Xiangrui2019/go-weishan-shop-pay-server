package services

import (
	"go-weishan-shop-pay-server/cache"
	"go-weishan-shop-pay-server/global"
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/modules"
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/utils"
	"net/http"
	"os"
	"strconv"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/qingwg/payjs/notify"
)

type ConfirmPayService struct {
}

func (service *ConfirmPayService) getOrderCache(msg notify.Message) (*global.OrderCache, error) {
	data := global.OrderCache{}

	orderjson, err := cache.CacheClient.Get(msg.Attach).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(orderjson), &data)
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
		SelfMention: data.SelfMention,
	}
}

func (service *ConfirmPayService) buildFee(data *global.OrderCache, to float64, fee float64) *models.Fee {
	return &models.Fee{
		TotalValue: data.BuyPrice,
		ToValue:    to,
		FeeValue:   fee,
	}
}

func (service *ConfirmPayService) createPayRecord(cachedata *global.OrderCache,
	to float64, fee float64) {
	tx := models.DB.Begin()

	result := tx.Create(service.buildOrder(cachedata))
	if result.Error != nil {
		tx.Rollback()
		panic(result.Error)
	}

	err := tx.Create(service.buildFee(cachedata, to, fee)).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func (service *ConfirmPayService) messageHandler(context *gin.Context, msg notify.Message) {
	data, err := service.getOrderCache(msg)

	if err != nil {
		panic(err)
	}

	feerate, err := strconv.ParseFloat(os.Getenv("FEE_RATE"), 64)

	if err != nil {
		panic(err)
	}

	to, fee := utils.CalcFee(data.BuyPrice, feerate)

	service.createPayRecord(data, to, fee)

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
