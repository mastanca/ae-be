package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Run("create debit Transaction", func(t *testing.T) {
			result := New(DebitTransaction, 100.0)
			assert.Equal(t, DebitTransaction, result.OperationType)
			assert.Equal(t, float64(100), result.Amount)
			assert.NotEmpty(t, result.Id)
			assert.NotNil(t, result.EffectiveDate)
		})
		t.Run("create credit Transaction", func(t *testing.T) {
			result := New(CreditTransaction, 500.0)
			assert.Equal(t, CreditTransaction, result.OperationType)
			assert.Equal(t, float64(500), result.Amount)
			assert.NotEmpty(t, result.Id)
			assert.NotNil(t, result.EffectiveDate)
		})
	})
}
