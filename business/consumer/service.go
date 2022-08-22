package consumer

import "context"

type service struct {
	repository Repositorer
}

func NewService(repository Repositorer) Servicer {
	return &service{repository}
}

func (s *service) InsertOne(data *Consumer) error {
	return s.repository.InsertOne(context.TODO(), data)
}

func (s *service) FindByID(id int32) (*Consumer, error) {
	return s.repository.FindByID(context.TODO(), id)
}
