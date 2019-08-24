package conf

import (
	"go-weishan-shop-pay-server/cache"
	"go-weishan-shop-pay-server/models"
	"go-weishan-shop-pay-server/modules"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
	modules.InitAllModules()
}
