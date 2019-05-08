package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	c "gitlab.com/mikrowezel/backend/granica/internal/config"

	"github.com/cenkalti/backoff"
	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	maxTries = 20
)

// Retry implements a retry mechanism in order to create a connected repo client.
func Retry(ctx context.Context, cfg *config.Config, logger log.Logger) chan *UserRepo {
	result := make(chan *UserRepo)

	bo := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(maxTries))

	go func() {
		defer close(result)

		for i := 0; i <= maxTries; i++ {

			conn, err := getConn(cfg)

			if err == nil {
				logger.Log("level", c.LogLevel.Info, "msg", "Mongo connection established")

				r := &UserRepo{
					conn: conn,
					coll: conn.Database("granica").Collection("users"),
				}
				result <- r
				return
			}

			logger.Log("level", c.LogLevel.Info, "msg", " Mongo connection error", "err", err.Error())

			// Backoff
			nb := bo.NextBackOff()
			if nb == backoff.Stop {
				result <- nil
				logger.Log("level", c.LogLevel.Info, "msg", "Mongo connection failed: Max number of tries reached.")
				bo.Reset()
				return
			}

			msg := fmt.Sprintf("Mongo connection failed: retrying in %s.", nb.String())
			logger.Log("level", c.LogLevel.Info, "msg", msg)
			time.Sleep(nb)
		}
	}()

	return result
}

// getConn returns a Mongo connected client.
func getConn(cfg *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("mongodb://%s:%d", cfg.Repo.MongoDB.Host, cfg.Repo.MongoDB.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, err
}
