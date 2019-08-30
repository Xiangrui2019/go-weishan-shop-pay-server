package tasks

import (
	"go-weishan-shop-pay-server/modules"
	"go-weishan-shop-pay-server/utils"
	"log"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs(locked bool) {
	Cron = cron.New()

	RegisterCronTasks()

	if !locked {
		if !modules.LockerModule.Lock("master", time.Minute*2) {
			Cron.AddFunc("@every 2m", CampaignMasterTask)
			Cron.Start()
			return
		}
	}

	Cron.AddFunc("@every 1m", ClifeMasterTask)

	for _, item := range modules.TasksModule {
		if item.Type == modules.TimeJob {
			d := item
			Cron.AddFunc(d.Time, func() {
				utils.PublishTask(d)
			})
		}
	}

	Cron.Start()

	log.Println("Cron Jobs started success.")
}
