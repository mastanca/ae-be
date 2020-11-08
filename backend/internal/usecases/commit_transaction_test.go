package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"
	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"
	"github.com/mastanca/accounting-notebook-be/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCommitTransactionImpl_Execute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		model := NewCommitTransactionModel(transaction.CreditTransaction, 500.0)
		t.Run("empty account", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			existingAccount := &account.Account{}
			repository.On("Get", mock.Anything).Return(existingAccount, nil)
			repository.On("Save", mock.Anything, mock.Anything).Return(nil)

			commitTransaction := NewCommitTransactionImpl(repository)
			committedTransaction, err := commitTransaction.Execute(context.TODO(), *model)

			assert.NoError(t, err)
			assert.NotNil(t, committedTransaction)
			assert.Len(t, existingAccount.Transactions, 1)
			assert.Equal(t, transaction.CreditTransaction, committedTransaction.OperationType)
			assert.Equal(t, float64(500), committedTransaction.Amount)
		})
		t.Run("preexistent account", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			existingAccount := &account.Account{Transactions: transaction.Transactions{transaction.New(transaction.DebitTransaction, 100.0)}}
			repository.On("Get", mock.Anything).Return(existingAccount, nil)
			repository.On("Save", mock.Anything, mock.Anything).Return(nil)

			commitTransaction := NewCommitTransactionImpl(repository)
			committedTransaction, err := commitTransaction.Execute(context.TODO(), *model)

			assert.NoError(t, err)
			assert.NotNil(t, committedTransaction)
			assert.Len(t, existingAccount.Transactions, 2)
			assert.Equal(t, transaction.CreditTransaction, committedTransaction.OperationType)
			assert.Equal(t, float64(500), committedTransaction.Amount)
		})
	})
	t.Run("Error", func(t *testing.T) {
		model := NewCommitTransactionModel(transaction.CreditTransaction, 500)
		t.Run("error getting account", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			repository.On("Get", mock.Anything).Return(nil, errors.New("fatal"))

			commitTransaction := NewCommitTransactionImpl(repository)
			committedTransaction, err := commitTransaction.Execute(context.TODO(), *model)

			assert.EqualError(t, err, "couldn't get account: fatal")
			assert.Nil(t, committedTransaction)
			repository.AssertNotCalled(t, "Save")
		})
		t.Run("nonexistent account", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			repository.On("Get", mock.Anything).Return(nil, nil)

			commitTransaction := NewCommitTransactionImpl(repository)
			committedTransaction, err := commitTransaction.Execute(context.TODO(), *model)

			assert.EqualError(t, err, "nonexistent account")
			assert.Nil(t, committedTransaction)
			repository.AssertNotCalled(t, "Save")
		})
		t.Run("error saving account", func(t *testing.T) {
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			existingAccount := &account.Account{}
			repository.On("Get", mock.Anything).Return(existingAccount, nil)
			repository.On("Save", mock.Anything, mock.Anything).Return(errors.New("fatal"))

			commitTransaction := NewCommitTransactionImpl(repository)
			committedTransaction, err := commitTransaction.Execute(context.TODO(), *model)

			assert.EqualError(t, err, "couldn't save account: fatal")
			assert.Nil(t, committedTransaction)
		})
		t.Run("invalid debit operation", func(t *testing.T) {
			model := NewCommitTransactionModel(transaction.DebitTransaction, 500)
			repository := new(mocks.Repository)
			defer repository.AssertExpectations(t)

			existingAccount := &account.Account{Transactions: transaction.Transactions{transaction.New(transaction.CreditTransaction, 200)}}
			repository.On("Get", mock.Anything).Return(existingAccount, nil)

			commitTransaction := NewCommitTransactionImpl(repository)
			committedTransaction, err := commitTransaction.Execute(context.TODO(), *model)

			assert.True(t, errors.Is(err, &account.InvalidTransactionError{}))
			assert.Nil(t, committedTransaction)
		})
	})
}
