package usecases

import (
	"context"

	"github.com/pkg/errors"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"

	"github.com/mastanca/accounting-notebook-be/internal/domain/transaction"
)

type CommitTransactionModel struct {
	OperationType transaction.TransactionType `json:"type"`
	Amount        uint64                      `json:"amount"`
}

func NewCommitTransactionModel(operationType transaction.TransactionType, amount uint64) *CommitTransactionModel {
	return &CommitTransactionModel{
		OperationType: operationType,
		Amount:        amount,
	}
}

type CommitTransaction interface {
	Execute(ctx context.Context, model CommitTransactionModel) (*transaction.Transaction, error)
}

type commitTransactionImpl struct {
	repository account.Repository
}

func (c commitTransactionImpl) Execute(ctx context.Context, model CommitTransactionModel) (*transaction.Transaction, error) {
	customerAccount, err := c.repository.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get account")
	}
	if customerAccount == nil {
		return nil, errors.New("nonexistent account")
	}
	transactionToCommit := transaction.New(model.OperationType, model.Amount)
	customerAccount.CommitTransaction(transactionToCommit)
	err = c.repository.Save(ctx, *customerAccount)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't save account")
	}
	return &transactionToCommit, nil
}

func NewCommitTransactionImpl(repository account.Repository) *commitTransactionImpl {
	return &commitTransactionImpl{repository: repository}
}

var _ CommitTransaction = (*commitTransactionImpl)(nil)
