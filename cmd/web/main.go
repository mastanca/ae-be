package main

import (
	"log"

	"github.com/mastanca/accounting-notebook-be/cmd/web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	pingHandler := handlers.NewPingHandlerImpl()

	router := gin.Default()
	router.GET("/ping", pingHandler.Handle)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("error initializing server")
	}
}
