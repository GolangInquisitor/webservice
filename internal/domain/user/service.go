package user

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, uuid string) error
}
