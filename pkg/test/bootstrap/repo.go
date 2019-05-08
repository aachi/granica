package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// getConn returns a Mongo connected client.
func getConn(cfg *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("mongodb://%s:%d", cfg.Repo.MongoDB.Host, cfg.Repo.MongoDB.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("[ERROR] MongoDB: Cannot create a client: %s", err.Error())
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("[ERROR] MongoDB: Cannot create a connection: %s", err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("[ERROR] MongoDB: connection ping error: %s", err.Error())
		return nil, err
	}

	log.Printf("[INFO] MongoDB: Connected!")

	return client, err
}
