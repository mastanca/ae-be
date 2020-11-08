package transaction

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string
type Transactions []transaction

const (
	DebitTransaction  TransactionType = "debit"
	CreditTransaction TransactionType = "credit"
)

type transaction struct {
	Id            string          `json:"id"`
	OperationType TransactionType `json:"type"` // type is a reserved word
	Amount        uint64          `json:"amount"`
	EffectiveDate time.Time       `json:"effective_date"`
}

func New(operationType TransactionType, amount uint64) *transaction {
	return &transaction{
		Id:            uuid.New().String(),
		OperationType: operationType,
		Amount:        amount,
		EffectiveDate: time.Now(),
	}
}