package main

import (
	"log"

	"github.com/mastanca/go-api-template/internal/usecases"

	"github.com/mastanca/go-api-template/cmd/web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	getUser := usecases.NewGetUserImpl()

	pingHandler := handlers.NewPingHandlerImpl()
	userLoginHandler := handlers.NewLoginHandlerImpl(getUser)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/ping", pingHandler.Handle)
		v1 := api.Group("/v1")
		{
			v1.POST("/login", userLoginHandler.Handle)
		}
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatal("error initializing server")
	}
}
