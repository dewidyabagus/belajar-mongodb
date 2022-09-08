package consumer

import (
	"belajar-mongodb/business"
	"belajar-mongodb/business/consumer"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timeOut time.Duration = 5 * time.Second
)

// Contoh penggunaan collection mongodb dengan contoh data consumers
type Consumer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Name      string             `bson:"name"`
	Address   string             `bson:"address"`
	CreatedAt int64              `bson:"created_at,omitempty"`
	UpdatedAt int64              `bson:"updated_at"`
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{db}
}

type Repository struct {
	db *mongo.Database
}

func (r *Repository) Collection() *mongo.Collection {
	return r.db.Collection("consumers")
}

func (r *Repository) InsertOne(ctx context.Context, data *consumer.Consumer) error {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	if _, err := r.Collection().InsertOne(ctxWT, r.newConsumer(data)); err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*consumer.Consumer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, business.ErrInvalidObjectID
	}

	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	result := new(Consumer)
	if err := r.Collection().FindOne(ctxWT, bson.M{"_id": objectID}).Decode(result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, business.ErrNotFound
		}
		return nil, err
	}

	return r.toBusinessConsumer(result), nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*consumer.Consumer, error) {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	result := new(Consumer)
	if err := r.Collection().FindOne(ctxWT, bson.M{"email": email}).Decode(result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, business.ErrNotFound
		}
		return nil, err
	}

	return r.toBusinessConsumer(result), nil
}

func (r *Repository) GetAll(ctx context.Context) ([]consumer.Consumer, error) {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut*2)
	defer cancel()

	cursor, err := r.Collection().Find(ctxWT, bson.M{})
	if err != nil {
		return nil, err
	}

	consumers := []Consumer{}

	ctxRead, cancelRead := context.WithTimeout(ctx, timeOut)
	defer cancelRead()

	if err := cursor.All(ctxRead, &consumers); err != nil {
		return nil, fmt.Errorf("reading all document: %w", err)
	}

	response := make([]consumer.Consumer, len(consumers))
	for i := range consumers {
		response[i] = *r.toBusinessConsumer(&consumers[i])
	}

	return response, nil
}

func (r *Repository) UpdateByID(ctx context.Context, id string, data *consumer.Consumer) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return business.ErrInvalidObjectID
	}

	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	result, err := r.Collection().UpdateByID(ctxWT, objectId, bson.M{"$set": r.updateConsumer(data)})
	if err != nil {
		return err

	} else if result.MatchedCount == 0 {
		return business.ErrNotFound
	}

	return nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return business.ErrInvalidObjectID
	}

	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	result, err := r.Collection().DeleteOne(ctxWT, bson.M{"_id": objectId})
	if err != nil {
		return err

	} else if result.DeletedCount == 0 {
		return business.ErrNotFound
	}

	return nil
}

func (r *Repository) newConsumer(data *consumer.Consumer) *Consumer {
	id, _ := primitive.ObjectIDFromHex(data.ID)

	return &Consumer{
		ID:        id,
		Email:     data.Email,
		Name:      data.Name,
		Address:   data.Address,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func (r *Repository) updateConsumer(c *consumer.Consumer) *Consumer {
	return &Consumer{
		Email:     c.Email,
		Name:      c.Name,
		Address:   c.Address,
		UpdatedAt: c.UpdatedAt,
	}
}

func (r *Repository) toBusinessConsumer(c *Consumer) *consumer.Consumer {
	return &consumer.Consumer{
		ID:        c.ID.Hex(),
		Email:     c.Email,
		Name:      c.Name,
		Address:   c.Address,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
