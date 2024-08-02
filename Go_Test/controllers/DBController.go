package controllers

import (
	"context"
	"fmt"
	"go_test/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() (*mongo.Client, error) { // Function Signature: The DB function now returns (*mongo.Client, error)
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.DBURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
