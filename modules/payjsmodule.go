package modules

import (
	"os"

	"github.com/qingwg/payjs"
)

var PayJSModule *payjs.PayJS

func InitPayJSModule() {
	PayJSModule = payjs.New(&payjs.Config{
		Key:       os.Getenv("MCH_KEY"),
		MchID:     os.Getenv("MCH_ID"),
		NotifyUrl: os.Getenv("PAY_CALLBACK"),
	})
}
