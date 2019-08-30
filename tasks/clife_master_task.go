package tasks

import (
	"go-weishan-shop-pay-server/cache"
	"go-weishan-shop-pay-server/global"
	"log"
	"time"
)

func ClifeMasterTask() {
	cache.CacheClient.Set(global.LockKey("master"), "true", time.Minute*2)
	log.Println("Continued life Success.")
}
