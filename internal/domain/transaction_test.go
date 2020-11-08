package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Run("create debit transaction", func(t *testing.T) {
			result := NewTransaction(debitTransaction, 100)
			assert.Equal(t, debitTransaction, result.OperationType)
			assert.Equal(t, uint64(100), result.Amount)
			assert.NotEmpty(t, result.Id)
			assert.NotNil(t, result.EffectiveDate)
		})
		t.Run("create credit transaction", func(t *testing.T) {
			result := NewTransaction(creditTransaction, 500)
			assert.Equal(t, creditTransaction, result.OperationType)
			assert.Equal(t, uint64(500), result.Amount)
			assert.NotEmpty(t, result.Id)
			assert.NotNil(t, result.EffectiveDate)
		})
	})
}
