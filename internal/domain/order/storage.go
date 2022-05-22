package order

import "context"

type Storage interface {
	Create(ctx context.Context, user_uuid string, order *Order) error
	Update(ctx context.Context, user_uuid string, order *Order) error
	Delete(ctx context.Context, uuid, orderId string) error
}
