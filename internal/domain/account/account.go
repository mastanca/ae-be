package account

import "github.com/mastanca/accounting-notebook-be/internal/domain/transaction"

type Account struct {
	Transactions transaction.Transactions
}
