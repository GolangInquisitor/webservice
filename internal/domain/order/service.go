package order

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, userUuid string, order *Order) error
	Update(ctx context.Context, userUuid string, order *Order) error
	Delete(ctx context.Context, uuid, orderId string) error
}
