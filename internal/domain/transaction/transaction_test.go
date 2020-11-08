package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Run("create debit transaction", func(t *testing.T) {
			result := New(DebitTransaction, 100)
			assert.Equal(t, DebitTransaction, result.OperationType)
			assert.Equal(t, uint64(100), result.Amount)
			assert.NotEmpty(t, result.Id)
			assert.NotNil(t, result.EffectiveDate)
		})
		t.Run("create credit transaction", func(t *testing.T) {
			result := New(CreditTransaction, 500)
			assert.Equal(t, CreditTransaction, result.OperationType)
			assert.Equal(t, uint64(500), result.Amount)
			assert.NotEmpty(t, result.Id)
			assert.NotNil(t, result.EffectiveDate)
		})
	})
}
