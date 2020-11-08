package transaction

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string
type Transactions []Transaction

const (
	DebitTransaction  TransactionType = "debit"
	CreditTransaction TransactionType = "credit"
)

type Transaction struct {
	Id            string          `json:"id"`
	OperationType TransactionType `json:"type"` // type is a reserved word
	Amount        float64         `json:"amount"`
	EffectiveDate time.Time       `json:"effective_date"`
}

func New(operationType TransactionType, amount float64) Transaction {
	return Transaction{
		Id:            uuid.New().String(),
		OperationType: operationType,
		Amount:        amount,
		EffectiveDate: time.Now(),
	}
}

func (t Transactions) FindById(id string) *Transaction {
	var result *Transaction
	for _, tt := range t {
		if tt.Id == id {
			result = &tt
		}
	}
	return result
}
