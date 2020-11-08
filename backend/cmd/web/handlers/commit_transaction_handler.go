package handlers

import (
	"errors"
	"net/http"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"
	"github.com/mastanca/accounting-notebook-be/internal/usecases"
)

type CommitTransactionHandler interface {
	Handle(c *gin.Context)
}

type commitTransactionHandlerImpl struct {
	commitTransaction usecases.CommitTransaction
}

func (ch commitTransactionHandlerImpl) Handle(c *gin.Context) {
	var newTransaction transaction.Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	committedTransaction, err := ch.commitTransaction.Execute(c, *usecases.NewCommitTransactionModel(newTransaction.OperationType, newTransaction.Amount))
	if err != nil {
		if errors.Is(err, &account.InvalidTransactionError{}) {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"reason": err.Error()})
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, committedTransaction)
}

func NewCommitTransactionHandler(commitTransaction usecases.CommitTransaction) *commitTransactionHandlerImpl {
	return &commitTransactionHandlerImpl{commitTransaction: commitTransaction}
}

var _ CommitTransactionHandler = (*commitTransactionHandlerImpl)(nil)
