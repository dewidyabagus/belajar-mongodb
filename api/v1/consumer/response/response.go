package response

import (
	"time"

	"belajar-mongodb/business/consumer"
)

type ResConsumer struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ResSingleConsumer(data *consumer.Consumer) *ResConsumer {
	return &ResConsumer{
		ID:        data.ID,
		Email:     data.Email,
		Name:      data.Name,
		Address:   data.Address,
		UpdatedAt: time.UnixMilli(data.UpdatedAt),
	}
}

func ResManyConsumer(data []consumer.Consumer) []ResConsumer {
	response := make([]ResConsumer, len(data))

	for i := range data {
		response[i] = *ResSingleConsumer(&data[i])
	}

	return response
}
