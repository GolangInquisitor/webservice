package service

import (
	usr "Scoltest/internal/domain/user"
	"Scoltest/pkg/loger"
	"context"
)

type Service struct {
	storage usr.Storage
	logger  *loger.Logger
}

func (s *Service) Create(ctx context.Context, user *usr.User) error {
	//TODO  Validate fields
	s.logger.Infof("new user Data: %#v", *user)
	if err := s.storage.Create(ctx, user); err != nil {
		s.logger.Error("Error create user struct: %#v . Message: %s", user, err.Error())
		return err
	}
	return nil
}
func (s *Service) Update(ctx context.Context, user *usr.User) error {
	//TODO  Validate fields
	s.logger.Infof("update user Data: %#v", *user)
	if err := s.storage.Update(ctx, user); err != nil {
		s.logger.Error("Error update user struct: %#v . Message: %s", user, err.Error())
		return err
	}
	return nil

}
func (s *Service) Delete(ctx context.Context, uuid string) error {
	//TODO  Validate fields
	s.logger.Infof("delete user uuid: %#v", uuid)
	if err := s.storage.Delete(ctx, uuid); err != nil {
		s.logger.Error("Error delete user uuid: %#v . Message: %s", uuid, err.Error())
		return err
	}
	return nil
}

func NewService(storage usr.Storage, logger *loger.Logger) *Service {

	return &Service{storage: storage, logger: logger}
}
