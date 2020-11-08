package usecases

import (
	"context"
	"testing"

	"github.com/pkg/errors"

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
		t.Run("existing transaction", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			repository.On("Get", mock.Anything).Return(&account.Account{Transactions: transaction.Transactions{transaction.Transaction{Id: id}}}, nil)

			getTransactionById := NewGetTransactionByIdImpl(repository)
			fetchedTransaction, err := getTransactionById.Execute(context.TODO(), id)

			assert.NoError(t, err)
			assert.NotNil(t, fetchedTransaction)
			assert.Equal(t, id, fetchedTransaction.Id)
		})
		t.Run("transaction not found", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			repository.On("Get", mock.Anything).Return(&account.Account{Transactions: transaction.Transactions{}}, nil)

			getTransactionById := NewGetTransactionByIdImpl(repository)
			fetchedTransaction, err := getTransactionById.Execute(context.TODO(), id)

			assert.NoError(t, err)
			assert.Nil(t, fetchedTransaction)
		})
	})
	t.Run("Error", func(t *testing.T) {
		t.Run("repository error", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			repository.On("Get", mock.Anything).Return(nil, errors.New("fatal"))

			getTransactionById := NewGetTransactionByIdImpl(repository)
			fetchedTransaction, err := getTransactionById.Execute(context.TODO(), id)

			assert.EqualError(t, err, "couldn't get account for the provided transaction: fatal")
			assert.Nil(t, fetchedTransaction)
		})
		t.Run("account not found", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			repository.On("Get", mock.Anything).Return(nil, nil)

			getTransactionById := NewGetTransactionByIdImpl(repository)
			fetchedTransaction, err := getTransactionById.Execute(context.TODO(), id)

			assert.EqualError(t, err, "nonexistent account")
			assert.Nil(t, fetchedTransaction)
		})
	})
}
