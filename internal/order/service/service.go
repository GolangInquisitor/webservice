package order

import (
	"Scoltest/internal/domain/order"
	"Scoltest/pkg/loger"
	"context"
)

type Service struct {
	storage order.Storage
	logger  *loger.Logger
}

func (s *Service) Create(ctx context.Context, userUuid string, order *order.Order) error {
	//TODO Validate userUuid and order
	s.logger.Infof("new user order Data: %#v", *order)
	if err := s.storage.Create(ctx, userUuid, order); err != nil {
		s.logger.Error("Error create user order struct: %#v . Message: %s", *order, err.Error())
		return err
	}
	return nil
}
func (s *Service) Update(ctx context.Context, userUuid string, order *order.Order) error {
	//TODO Validate user_uuid and order
	s.logger.Infof("update user order Data: %#v", *order)
	if err := s.storage.Update(ctx, userUuid, order); err != nil {
		s.logger.Error("Error update user order struct: %#v . Message: %s", *order, err.Error())
		return err
	}
	return nil
}
func (s *Service) Delete(ctx context.Context, uuid, orderId string) error {
	//TODO Validate user_uuid
	s.logger.Infof("delete order uuid: %s , order id: %s", uuid, orderId)
	if err := s.storage.Delete(ctx, uuid, orderId); err != nil {
		s.logger.Error("Error delete user uuid: %#v . Message: %s", uuid, err.Error())
		return err
	}
	return nil
}

func NewService(storage order.Storage, logger *loger.Logger) *Service {

	return &Service{storage: storage, logger: logger}
}
