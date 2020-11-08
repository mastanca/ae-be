package account

import "github.com/mastanca/accounting-notebook-be/internal/domain/transaction"

type Account struct {
	Transactions transaction.Transactions `json:"transactions"`
}

func (a *Account) CommitTransaction(transaction transaction.Transaction) {
	a.Transactions = append(a.Transactions, transaction)
}

func (a Account) GetBalance() float64 {
	var result float64
	for _, t := range a.Transactions {
		switch t.OperationType {
		case transaction.DebitTransaction:
			result -= t.Amount
		case transaction.CreditTransaction:
			result += t.Amount
		}
	}
	return result
}

type InvalidTransactionError struct {
}

func (e *InvalidTransactionError) Error() string { return "insufficient funds" }
