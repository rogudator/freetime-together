package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context, cfg ConfigMongo) (*mongo.Client, error) {
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI(ConfiMongogURI(cfg)))
	if err != nil {
		return nil, err
	}
	// check the connection
	err = connect.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return connect, nil
}

func ConfiMongogURI(cfg ConfigMongo) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	return uri
}

type ConfigMongo struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}
