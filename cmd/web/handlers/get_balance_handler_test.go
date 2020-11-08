package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"
	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"
	"github.com/mastanca/accounting-notebook-be/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

type balanceResponse struct {
	Balance float64 `json:"balance"`
}

func TestGetBalanceHandlerImpl_Handle(t *testing.T) {
	getAccount := new(mocks.GetAccount)
	defer getAccount.AssertExpectations(t)

	getAccount.On("Execute", mock.Anything).Return(&account.Account{Transactions: transaction.Transactions{transaction.New(transaction.CreditTransaction, 100)}}, nil)

	handler := NewGetBalanceHandlerImpl(getAccount)

	r := gin.Default()
	r.GET("/", handler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	var responseBody balanceResponse
	_ = json.Unmarshal(w.Body.Bytes(), &responseBody)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, float64(100), responseBody.Balance)
}
