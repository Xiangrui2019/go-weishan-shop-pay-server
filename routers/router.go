package routers

import (
	"go-weishan-shop-pay-server/api"
	"go-weishan-shop-pay-server/middlewares"
	"go-weishan-shop-pay-server/tasks"
	"go-weishan-shop-pay-server/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)

		v1.POST("/pay/create", api.CreatePay)
		v1.POST("/pay/update", api.ConfirmPay)
	}

	task := router.Group("/tasks/v1")
	{
		task.GET("/compute_report", func(context *gin.Context) {
			utils.RunTask(context, tasks.ComputeReportTask)
		})
	}

	web := router.Group("/web/v1")
	{
		web.GET("/order/put/:id", api.PublishOrder)
		web.GET("/order/check/:id", api.CheckPublishOrder)
	}

	return router
}
