package tasks

import (
	"go-weishan-shop-pay-server/modules"
)

func RegisterCronTasks() {
	modules.ClearTimedJob()
	modules.AddTimedJob("@every 1m", TimeTask)
	modules.AddTimedJob("@every 2m", DeltaTask)
	modules.AddAyncJob(TimeTask1)
}
