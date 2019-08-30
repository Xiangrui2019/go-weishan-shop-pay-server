package routers

import (
	"go-weishan-shop-pay-server/api"
	"go-weishan-shop-pay-server/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)

		v1.GET("/order", api.ListOrder)
		v1.GET("/order/non_publish", api.ListNonPubOrder)

		v1.POST("/pay/create", api.CreatePay)
		v1.POST("/pay/confirm", api.ConfirmPay)
	}

	web := router.Group("/web/v1")
	{
		web.GET("/order/publish/:id", api.PublishOrder)
	}

	return router
}
