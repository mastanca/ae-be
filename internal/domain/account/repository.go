package account

import (
	"context"
	"sync"
)

type Repository interface {
	Get(ctx context.Context) (*Account, error)
	Save(ctx context.Context, account Account) error
}

type inMemoryRepositoryImpl struct {
	m       sync.RWMutex
	account Account
}

func NewInMemoryRepositoryImpl() *inMemoryRepositoryImpl {
	return &inMemoryRepositoryImpl{}
}

func (i *inMemoryRepositoryImpl) Get(ctx context.Context) (*Account, error) {
	i.m.RLock()
	defer i.m.RUnlock()
	return &i.account, nil
}

func (i *inMemoryRepositoryImpl) Save(ctx context.Context, account Account) error {
	i.m.Lock()
	defer i.m.Unlock()
	i.account = account
	return nil
}

var _ Repository = (*inMemoryRepositoryImpl)(nil)
