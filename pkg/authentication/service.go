/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"gitlab.com/mikrowezel/backend/granica/pkg/authentication/repo"
	m "gitlab.com/mikrowezel/backend/granica/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

var (
// jwtKeys *config.JWTKeys
)

// ErrServer is a general server error
var ErrServer = errors.New("Internal server error")

// GranicaService provides authentication and authorization services
type GranicaService interface {
	SignUp(username, password, email, emailConfirmation, tenantID string) (*m.User, error)
	Cancel(username, password, tenantID string) error
	SignIn(username, password, tenantID string) (*m.User, error)
	SignOut(username, tenantID string) error
	Create(username, password, email, tenantID string) (*m.User, error)
	Update(username, password, passwordConfirmation,
		email, emailConfirmation, description,
		givenName, middleNames, familyName, tenantID string) error
	Remove(username, email, tenantID string) error
	Logger() log.Logger
}

type granicaService struct {
	name    string
	ctx     context.Context
	cfg     *config.Config
	logger  log.Logger
	repo    repo.UserRepo
	code    int
	message string
	err     error
}

// Interface implementation

// SignUp signs a user up using username, password, email and tenant.
func (gs granicaService) SignUp(username, password, email, emailConfirmation, tenantID string) (*m.User, error) {
	user := m.User{
		Username: username,
		Password: password,
		Email:    email,
		TenantID: tenantID,
	}

	_, err := gs.repo.Insert(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Cancel lets the user cancel his/her account.
func (gs granicaService) Cancel(username, password, tenantID string) error {
	user, err := gs.repo.GetByUsernameAndTenant(username, tenantID)
	if err != nil {
		return err
	}

	if !passwordMatches(user.PasswordDigest, password) {
		return errors.New("password doesn't match")
	}

	err = gs.repo.Delete(user.ID)
	if err != nil {
		return err
	}

	return nil
}

// SignIn lets a user sign in providing username/email, password and tenant.
func (gs granicaService) SignIn(username, password, tenantID string) (*m.User, error) {
	user, err := gs.repo.GetByUsernameAndTenant(username, tenantID)
	if err != nil {
		return nil, err
	}

	if !passwordMatches(user.PasswordDigest, password) {
		return nil, errors.New("password doesn't match")
	}

	return user, nil
}

// SignOut lets a user sign out.
func (gs granicaService) SignOut(username, tenantID string) error {
	// TODO: Close session implementation.
	return nil
}

// Create lets the system administrator create a user.
func (gs granicaService) Create(username, password, email, tenantID string) (*m.User, error) {
	user := m.User{
		Username: username,
		Password: password,
		Email:    email,
		TenantID: tenantID,
	}

	_, err := gs.repo.Insert(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update lets the system administrator user a user.
func (gs granicaService) Update(username, password, passwordConfirmation,
	email, emailConfirmation, description,
	givenName, middleNames, familyName, tenantID string) error {
	user, err := gs.repo.GetByUsernameAndTenant(username, tenantID)
	if err != nil {
		return err
	}

	if !passwordMatches(user.PasswordDigest, password) {
		return errors.New("password doesn't match")
	}

	// Copy values from current but only update allowed fields.
	// TODO: Validations.
	// TODO: Audit fields auto update.
	u := user
	u.Password = password
	u.Email = email
	u.Description = description
	u.GivenName = givenName
	u.MiddleNames = middleNames
	u.FamilyName = familyName
	u.SetUpdateValues(u.ID)

	err = gs.repo.Update(u)
	if err != nil {
		return err
	}

	return nil
}

// Create lets system administrator remove a user providing username/email and tenant.
func (gs granicaService) Remove(username, email, tenantID string) error {
	// tx := gs.newTx()
	// rs := makeRelayService(tx, gs.logger)
	// user, err := rs.Remove(username, email, tenantID)
	// if tx.Commit() != nil {
	// 	return nil, err
	// }
	// if err != nil {
	// 	return nil, err
	// }
	// return user, nil
	return nil
}

func (gs granicaService) Logger() log.Logger {
	return gs.logger
}

func passwordMatches(digest, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(password))
	if err != nil {
		return false
	}
	return true
}
