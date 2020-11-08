package main

import (
	"log"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"
	"github.com/mastanca/accounting-notebook-be/internal/usecases"

	"github.com/mastanca/accounting-notebook-be/cmd/web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	accountRepository := account.NewInMemoryRepositoryImpl()

	getAccount := usecases.NewGetAccountImpl(accountRepository)
	commitTransaction := usecases.NewCommitTransactionImpl(accountRepository)
	getTransactionById := usecases.NewGetTransactionByIdImpl(accountRepository)

	pingHandler := handlers.NewPingHandlerImpl()
	getBalanceHandler := handlers.NewGetBalanceHandlerImpl(getAccount)
	getTransactionsHandler := handlers.NewGetTransactionsHandlerImpl(getAccount)
	commitTransactionsHandler := handlers.NewCommitTransactionHandler(commitTransaction)
	getTransactionByIdHandler := handlers.NewGetTransactionByIdHandlerImpl(getTransactionById)

	router := gin.Default()
	router.GET("/", getBalanceHandler.Handle)
	router.GET("/ping", pingHandler.Handle)
	transactions := router.Group("/transactions")
	{
		transactions.GET("", getTransactionsHandler.Handle)
		transactions.POST("", commitTransactionsHandler.Handle)
		transactions.GET("/:id", getTransactionByIdHandler.Handle)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("error initializing server")
	}
}
