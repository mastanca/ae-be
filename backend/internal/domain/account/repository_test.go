package account

import (
	"context"
	"testing"

	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepositoryImpl_Get(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repository := NewInMemoryRepositoryImpl()
		account, err := repository.Get(context.TODO())
		assert.NoError(t, err)
		assert.IsType(t, account, &Account{})
	})
}

func TestInMemoryRepositoryImpl_Save(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repository := NewInMemoryRepositoryImpl()
		transactions := transaction.Transactions{transaction.New(transaction.DebitTransaction, 100)}
		err := repository.Save(context.TODO(), Account{Transactions: transactions})
		assert.NoError(t, err)
	})
}
