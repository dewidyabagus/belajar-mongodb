package consumer

import (
	"belajar-mongodb/business/consumer"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timeOut time.Duration = 5 * time.Second
)

type Consumer struct {
	ID      int32  `bson:"_id"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
}

type Repository struct {
	db         *mongo.Database
	collection string
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{db, "consumers"}
}

func (r *Repository) Collection() *mongo.Collection {
	return r.db.Collection(r.collection)
}

func (r *Repository) InsertOne(ctx context.Context, data *consumer.Consumer) error {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	result, err := r.Collection().InsertOne(ctxWT, Consumer{ID: data.ID, Name: data.Name, Address: data.Address})
	if err != nil {
		return err
	}
	fmt.Println("Result insert one:", result)

	return nil
}

func (r *Repository) FindByID(ctx context.Context, id int32) (*consumer.Consumer, error) {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut*2) // All process must be complate 10 s
	defer cancel()

	result := new(Consumer)
	if err := r.Collection().FindOne(ctxWT, bson.M{"_id": id}).Decode(result); err != nil {
		return nil, err
	}

	return &consumer.Consumer{
		ID:      result.ID,
		Name:    result.Name,
		Address: result.Address,
	}, nil
}
