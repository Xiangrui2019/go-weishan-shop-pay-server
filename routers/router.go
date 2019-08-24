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

		v1.GET("/order/put/:id", api.PublishOrder)
		v1.GET("/order", api.ListOrder)
		v1.GET("/order/check/:id", api.CheckPublishOrder)
		v1.POST("/pay/create", api.CreateOrder)
		v1.POST("/order/update", api.FinishOrder)
	}

	task := router.Group("/tasks")
	{
		task.GET("/compute_report", func(context *gin.Context) { utils.RunTask(context, tasks.ComputeReportTask) })
	}

	return router
}
