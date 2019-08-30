package tasks

import (
	"go-weishan-shop-pay-server/modules"
)

func RegisterCronTasks() {
	modules.ClearTimedJob()
	modules.AddAyncJob(ConfirmTask)
}
