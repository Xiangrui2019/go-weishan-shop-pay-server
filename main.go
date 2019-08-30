package main

import (
	_ "go-weishan-shop-pay-server/conf"
	"go-weishan-shop-pay-server/routers"
	"log"
	"go-weishan-shop-pay-server/protocol/http"
	"os"
)

func main() {
	router := routers.NewRouter()
	server := http.NewHttpProtocol(router)

	err := server.Start(os.Getenv("ADDR"))

	if err != nil {
		log.Fatal(err)
	}
}
