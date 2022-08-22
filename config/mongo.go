package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(ctx context.Context, config MongoConfig) (*mongo.Database, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?maxPoolSize=%s",
		config.Username, config.Password, config.Host, config.Port, config.MaxPoolSize,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error connect: %w", err)
	}

	return client.Database(config.Database), nil
}
