package usecases

import (
	"context"

	"github.com/pkg/errors"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"

	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"
)

type GetTransactionById interface {
	Execute(ctx context.Context, id string) (*transaction.Transaction, error)
}

type getTransactionByIdImpl struct {
	repository account.Repository
}

func (g getTransactionByIdImpl) Execute(ctx context.Context, id string) (*transaction.Transaction, error) {
	existingAccount, err := g.repository.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get account for the provided transaction")
	}
	if existingAccount == nil {
		return nil, errors.New("nonexistent account")
	}
	return existingAccount.Transactions.FindById(id), nil
}

func NewGetTransactionByIdImpl(repository account.Repository) *getTransactionByIdImpl {
	return &getTransactionByIdImpl{repository: repository}
}

var _ GetTransactionById = (*getTransactionByIdImpl)(nil)
