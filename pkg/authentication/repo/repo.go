package repo

import (
	// c "gitlab.com/mikrowezel/backend/granica/internal/config"
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"gitlab.com/mikrowezel/backend/granica/pkg/authentication/repo/mongodb"
	m "gitlab.com/mikrowezel/backend/granica/pkg/models"
)

// UserRepo interface
type UserRepo interface {
	Insert(*m.User) (id interface{}, err error)
	Get(id interface{}) (*m.User, error)
	GetAll() ([]m.User, error)
	GetByUsernameAndTenant(username, tenantID string) (*m.User, error)
	Update(*m.User) error
	Delete(id interface{}) error
}

// NewRepo makes a new user repo.
func NewRepo(ctx context.Context, cfg *config.Config, logger log.Logger) (UserRepo, error) {

	if cfg.Repo.Type == "mongodb" {
		return mongodb.NewRepo(ctx, cfg, logger)

	} else if cfg.Repo.Type == "dgraph" {
		// return dgraph.NewRepo(ctx, cfg, logger)
		return nil, errors.New("not a valid repo type")
	}

	return nil, errors.New("not a valid repo type")
}
