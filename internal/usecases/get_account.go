package usecases

import (
	"context"

	"github.com/mastanca/accounting-notebook-be/internal/domain/account"
)

type GetAccount interface {
	Execute(ctx context.Context) (*account.Account, error)
}

type getAccountImpl struct {
	repository account.Repository
}

func (g getAccountImpl) Execute(ctx context.Context) (*account.Account, error) {
	return g.repository.Get(ctx)
}

func NewGetAccountImpl(repository account.Repository) *getAccountImpl {
	return &getAccountImpl{repository: repository}
}

var _ GetAccount = (*getAccountImpl)(nil)
