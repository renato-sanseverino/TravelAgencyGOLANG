package repository

import (
	"context"
)

// usar generics na interface (golang 1.18)
type IRepository[T any] interface {
	Insert(context.Context, T) error
	GetByID(context.Context, int) (*T, error)
	Delete(context.Context, int) error
	Patch(context.Context, int, T) error
}
