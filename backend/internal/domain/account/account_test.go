package account

import (
	"testing"

	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"
	"github.com/stretchr/testify/assert"
)

func TestAccount_CommitTransaction(t *testing.T) {
	t.Run("should append new transaction", func(t *testing.T) {
		newAccount := Account{}
		transactionToCommit := transaction.New(transaction.DebitTransaction, 100.0)
		newAccount.CommitTransaction(transactionToCommit)
		assert.NotNil(t, newAccount.Transactions)
		assert.Len(t, newAccount.Transactions, 1)
		assert.Equal(t, transactionToCommit.OperationType, newAccount.Transactions[0].OperationType)
		assert.Equal(t, transactionToCommit.Amount, newAccount.Transactions[0].Amount)
	})
}

func TestAccount_GetBalance(t *testing.T) {
	transaction1 := transaction.New(transaction.CreditTransaction, 200.0)
	transaction2 := transaction.New(transaction.DebitTransaction, 100.0)
	newAccount := &Account{Transactions: transaction.Transactions{transaction1, transaction2}}
	assert.Equal(t, 100.0, newAccount.GetBalance())
}
