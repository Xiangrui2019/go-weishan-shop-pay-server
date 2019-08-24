package main

import (
	_ "go-weishan-shop-pay-server/conf"
	"go-weishan-shop-pay-server/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routers.NewRouter()

	httpServer := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: router,
	}

	log.Printf("Server started on %s", os.Getenv("ADDR"))
	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
