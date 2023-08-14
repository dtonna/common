package repository

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWork, error)
	Commit() error
	Rollback() error
}
