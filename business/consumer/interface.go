package consumer

import "context"

type Servicer interface {
	InsertOne(ctx context.Context, data *NewConsumer) error

	GetAll(ctx context.Context) ([]Consumer, error)
	FindByID(ctx context.Context, id string) (*Consumer, error)

	UpdateByID(ctx context.Context, id string, data *UpConsumer) error

	DeleteByID(ctx context.Context, id string) error
}

type Repositorer interface {
	InsertOne(ctx context.Context, data *Consumer) error

	GetAll(ctx context.Context) ([]Consumer, error)
	FindByID(ctx context.Context, id string) (*Consumer, error)
	FindByEmail(ctx context.Context, email string) (*Consumer, error)

	UpdateByID(ctx context.Context, id string, data *Consumer) error

	DeleteByID(ctx context.Context, id string) error
}
