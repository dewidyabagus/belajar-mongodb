package consumer

import "context"

type Servicer interface {
	InsertOne(data *Consumer) error

	FindByID(id int32) (*Consumer, error)
}

type Repositorer interface {
	InsertOne(ctx context.Context, data *Consumer) error

	FindByID(ctx context.Context, id int32) (*Consumer, error)
}
