package consumer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewConsumer struct {
	Email   string `validate:"email"`
	Name    string `validate:"required"`
	Address string `validate:"required"`
}

func (n *NewConsumer) ToConsumer() *Consumer {
	return &Consumer{
		ID:        primitive.NewObjectID().Hex(),
		Email:     n.Email,
		Name:      n.Name,
		Address:   n.Address,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
}

type UpConsumer struct {
	Email   string `validate:"email"`
	Name    string `validate:"required"`
	Address string `validate:"required"`
}

func (u *UpConsumer) ToConsumer() *Consumer {
	return &Consumer{
		Email:     u.Email,
		Name:      u.Name,
		Address:   u.Address,
		UpdatedAt: time.Now().UnixMilli(),
	}
}

type Consumer struct {
	ID        string
	Email     string
	Name      string
	Address   string
	CreatedAt int64
	UpdatedAt int64
}
