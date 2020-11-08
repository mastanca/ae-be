package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mastanca/accounting-notebook-be/internal/usecases"
)

type GetTransactionByIdHandler interface {
	Handle(c *gin.Context)
}

type getTransactionByIdHandlerImpl struct {
	getTransactionById usecases.GetTransactionById
}

func (g getTransactionByIdHandlerImpl) Handle(c *gin.Context) {
	id := c.Param("id")
	fetchedTransaction, err := g.getTransactionById.Execute(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if fetchedTransaction == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"reason": "transaction not found"})
		return
	}
	c.JSON(http.StatusOK, fetchedTransaction)
}

func NewGetTransactionByIdHandlerImpl(getTransactionById usecases.GetTransactionById) *getTransactionByIdHandlerImpl {
	return &getTransactionByIdHandlerImpl{getTransactionById: getTransactionById}
}

var _ GetTransactionByIdHandler = (*getTransactionByIdHandlerImpl)(nil)
