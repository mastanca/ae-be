package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/accounting-notebook-be/internal/usecases"
)

type GetBalanceHandler interface {
	Handle(c *gin.Context)
}

type getBalanceHandlerImpl struct {
	getAccount usecases.GetAccount
}

func (g getBalanceHandlerImpl) Handle(c *gin.Context) {
	account, err := g.getAccount.Execute(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if account == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"reason": "nonexistent account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": account.GetBalance()})
}

func NewGetBalanceHandlerImpl(getAccount usecases.GetAccount) *getBalanceHandlerImpl {
	return &getBalanceHandlerImpl{getAccount: getAccount}
}

var _ GetBalanceHandler = (*getBalanceHandlerImpl)(nil)
