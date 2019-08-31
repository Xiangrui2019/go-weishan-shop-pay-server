package services

import (
	"encoding/json"
	"go-weishan-shop-pay-server/cache"
	"go-weishan-shop-pay-server/global"
	"go-weishan-shop-pay-server/modules"
	"go-weishan-shop-pay-server/serializer"
	"go-weishan-shop-pay-server/tasks"
	"go-weishan-shop-pay-server/utils"
	"net/http"

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

	_, err = cache.CacheClient.Del(msg.Attach).Result()
	if err != nil {
		panic(err)
	}

	return &data, nil
}

func (service *ConfirmPayService) messageHandler(msg notify.Message) {
	data, err := service.getOrderCache(msg)

	ds, _ := json.Marshal(&data)

	err = utils.RunAsyncTask(tasks.ConfirmTask, string(ds))

	if err != nil {
		panic(err)
	}
}

func (service *ConfirmPayService) Finish(context *gin.Context) *serializer.Response {
	payNotify := modules.PayJSModule.GetNotify(context.Request, context.Writer)

	payNotify.SetMessageHandler(func(msg notify.Message) {
		service.messageHandler(msg)
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
