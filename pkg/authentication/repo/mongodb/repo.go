package mongodb

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// c "gitlab.com/mikrowezel/backend/granica/internal/config"
	m "gitlab.com/mikrowezel/backend/granica/pkg/models"
)

// UserRepo is a Mongo implementation of UserRepo interface.
type UserRepo struct {
	ctx  context.Context
	conn *mongo.Client
	coll *mongo.Collection
}

// NewRepo makes a new user repo.
func NewRepo(ctx context.Context, cfg *config.Config, logger log.Logger) (*UserRepo, error) {
	// conn, err := getConn(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	// return &UserRepo{
	// 	conn: conn,
	// 	coll: conn.Database("granica").Collection("users"),
	// }, nil
	return <-Retry(ctx, cfg, logger), nil

}

// Insert a user in UserRepo.
func (r *UserRepo) Insert(user *m.User) (id interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

// GetAll users from repo.
func (r *UserRepo) GetAll() ([]m.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []m.User

	cur, err := r.coll.Find(ctx, bson.D{})
	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		u := m.User{}
		err = cur.Decode(&u)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Get a users from repo by its ID.
func (r *UserRepo) Get(id interface{}) (*m.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user m.User

	filter := bson.M{"ID": id}
	err := r.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Get a users from repo by its username and tenant.
func (r *UserRepo) GetByUsernameAndTenant(username, tenantID string) (*m.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user m.User

	filter := bson.M{"Username": username, "TenantID": tenantID}
	err := r.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update a user in UserRepo.
func (r *UserRepo) Update(user *m.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"ID": user.ID}
	_, err := r.coll.UpdateOne(ctx, filter, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user from repo.
func (r *UserRepo) Delete(id interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"ID": id}
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
