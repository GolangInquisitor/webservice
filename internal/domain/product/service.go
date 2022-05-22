package product

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, uuid string) error
}
