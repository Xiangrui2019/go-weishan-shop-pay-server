package tasks

import (
	"go-weishan-shop-pay-server/modules"
	"log"
	"time"
)

func CampaignMasterTask() {
	if modules.LockerModule.Lock("master", time.Minute*2) {
		StartCronJobs(true)
		log.Println("Campaign master Success.")
	}
}
