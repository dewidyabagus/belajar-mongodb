package request

import "belajar-mongodb/business/consumer"

type NewConsumer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (n *NewConsumer) ToBusinessConsumer() *consumer.NewConsumer {
	return &consumer.NewConsumer{
		Email:   n.Email,
		Name:    n.Name,
		Address: n.Address,
	}
}

type UpConsumer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (u *UpConsumer) ToBusinessConsumer() *consumer.UpConsumer {
	return &consumer.UpConsumer{
		Email:   u.Email,
		Name:    u.Name,
		Address: u.Address,
	}
}
