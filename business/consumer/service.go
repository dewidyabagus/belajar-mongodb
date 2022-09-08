package consumer

import (
	"belajar-mongodb/business"
	"belajar-mongodb/pkg/validator"
	"context"
	"errors"
)

type service struct {
	repository Repositorer
}

func NewService(repository Repositorer) Servicer {
	return &service{repository}
}

func (s *service) InsertOne(ctx context.Context, data *NewConsumer) error {
	if err := validator.ValidateStruct(data); err != nil {
		return business.ErrorNotValid(err.Error())
	}

	if err := s.validateEmail(ctx, data.Email); err != nil {
		return err
	}

	return s.repository.InsertOne(ctx, data.ToConsumer())
}

func (s *service) FindByID(ctx context.Context, id string) (*Consumer, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]Consumer, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) UpdateByID(ctx context.Context, id string, data *UpConsumer) error {
	if err := validator.ValidateStruct(data); err != nil {
		return business.ErrorNotValid(err.Error())
	}

	currentConsumer, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if currentConsumer.Email != data.Email {
		if err := s.validateEmail(ctx, data.Email); err != nil {
			return err
		}
	}

	return s.repository.UpdateByID(ctx, id, data.ToConsumer())
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	return s.repository.DeleteByID(ctx, id)
}

func (s *service) validateEmail(ctx context.Context, email string) error {
	if _, err := s.repository.FindByEmail(ctx, email); err == nil {
		return business.ErrorConflict("email already exists")

	} else if !errors.Is(err, business.ErrNotFound) {
		return err

	}
	return nil
}
