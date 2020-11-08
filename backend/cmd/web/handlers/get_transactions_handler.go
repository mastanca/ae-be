package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/accounting-notebook-be/internal/usecases"
)

type GetTransactionsHandler interface {
	Handle(c *gin.Context)
}

type getTransactionsHandlerImpl struct {
	getAccount usecases.GetAccount
}

func (g getTransactionsHandlerImpl) Handle(c *gin.Context) {
	account, err := g.getAccount.Execute(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if account == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"reason": "nonexistent account"})
		return
	}
	c.JSON(http.StatusOK, account.Transactions)
}

func NewGetTransactionsHandlerImpl(getAccount usecases.GetAccount) *getTransactionsHandlerImpl {
	return &getTransactionsHandlerImpl{getAccount: getAccount}
}

var _ GetTransactionsHandler = (*getTransactionsHandlerImpl)(nil)
