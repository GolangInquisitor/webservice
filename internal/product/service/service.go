package product

import (
	"Scoltest/internal/domain/product"

	"Scoltest/pkg/loger"
	"context"
)

type Service struct {
	storage product.Storage
	logger  *loger.Logger
}

func (s *Service) Create(ctx context.Context, product *product.Product) error {
	//TODO  Validate product fields
	s.logger.Infof("new product Data: %#v", *product)
	if err := s.storage.Create(ctx, product); err != nil {
		s.logger.Error("Error create product struct: %#v . Message: %s", product, err.Error())
		return err
	}
	return nil
}
func (s *Service) Update(ctx context.Context, product *product.Product) error {
	//TODO  Validate product fields
	s.logger.Infof("new product Data: %#v", *product)
	if err := s.storage.Update(ctx, product); err != nil {
		s.logger.Error("Error create product struct: %#v . Message: %s", product, err.Error())
		return err
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, uuid string) error {
	s.logger.Infof("delete product uuid: %#v", uuid)
	if err := s.storage.Delete(ctx, uuid); err != nil {
		s.logger.Error("Error delete product uuid: %#v . Message: %s", uuid, err.Error())
		return err
	}
	return nil
}

func NewService(storage product.Storage, logger *loger.Logger) *Service {

	return &Service{storage: storage, logger: logger}
}
