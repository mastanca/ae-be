package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"

	"github.com/google/uuid"
	"github.com/mastanca/accounting-notebook-be/internal/domain/account"
	"github.com/mastanca/accounting-notebook-be/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGetTransactionByIdImpl_Execute(t *testing.T) {
	id := uuid.New().String()
	t.Run("Success", func(t *testing.T) {
		repository := new(mocks.Repository)
		defer repository.AssertExpectations(t)

		repository.On("Get", mock.Anything).Return(&account.Account{Transactions: transaction.Transactions{transaction.Transaction{Id: id}}}, nil)

		getTransactionById := NewGetTransactionByIdImpl(repository)
		fetchedTransaction, err := getTransactionById.Execute(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, fetchedTransaction)
		assert.Equal(t, id, fetchedTransaction.Id)
	})
}
