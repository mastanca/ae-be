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
		if fetchedTransaction == nil {
			_ = c.AbortWithError(http.StatusNotFound, err)
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, fetchedTransaction)
}

func NewGetTransactionByIdHandlerImpl(getTransactionById usecases.GetTransactionById) *getTransactionByIdHandlerImpl {
	return &getTransactionByIdHandlerImpl{getTransactionById: getTransactionById}
}

var _ GetTransactionByIdHandler = (*getTransactionByIdHandlerImpl)(nil)
