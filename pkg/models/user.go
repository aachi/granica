/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userModelName = "user"
)

// User model struct
type User struct {
	MongoID `bson:",inline"`
	Identification
	TenantID             string    `bson:"tenant_id" json:"tenantID"`
	Username             string    `bson:"username" json:"username"`
	Password             string    `bson:"-" json:"password"`
	PasswordConfirmation string    `bson:"-" json:"passwordConfirmation"`
	PasswordDigest       string    `bson:"password_digest" json:"-"`
	Email                string    `bson:"email" json:"email"`
	EmailConfirmation    string    `bson:"-" json:"emailConfirmation"`
	Description          string    `bson:"description" json:"description"`
	GivenName            string    `bson:"given_name" json:"givenName"`
	MiddleNames          string    `bson:"middle_names" json:"middleNames"`
	FamilyName           string    `bson:"family_name" json:"familyNames"`
	ContextID            string    `bson:"context_id" json:"contextID"`
	LastKnownIP          string    `bson:"last_known_ip" json:"lastKnownIP"`
	LastKnownLocation    Geometry  `bson:"last_known_location" json:"lastKnownLocation"`
	Geohash              string    `bson:"geohash" json:"geohash"`
	StartsAt             time.Time `bson:"starts_at" json:"startsAt"`
	EndsAt               time.Time `bson:"ends_at" json:"endsAt"`
	IsActive             bool      `bson:"is_active" json:"isActive"`
	IsLogicalDeleted     bool      `bson:"is_logical_deleted" json:"isLogicalDeleted"`
	Audit
	Persist
	// A user can only own a single 'individual' account
	// but multiple of 'organization' type.
	// Accounts *Account `has_many:"accounts" order_by:"account_type asc"`
	// Tenure   *User    `has_many_to:"tenures"`
}

// SetCreateValues sets creation values
// FIX: This method must be extracted
// to a generic interface embedded in all models.
func (user *User) SetCreateValues(tenantID string, createdBy ...uuid.UUID) {
	user.GenAndSetID()
	user.GenAndSetNameKey()
	user.UpdateSlug()
	user.UpdatePasswordDigest()
	user.SetTenant(tenantID)
	user.SetCreatedBy(user.idOrSelf(createdBy...))
	user.SetUpdatedBy(user.idOrSelf(createdBy...))
	user.SetCreatedAt()
	user.SetUpdatedAt()
}

// SetUpdateValues sets creation values
// FIX: This method must be extracted
// to a generic interface embedded in all models.
func (user *User) SetUpdateValues(updatedBy ...uuid.UUID) {
	user.UpdateSlug()
	user.UpdatePasswordDigest()
	user.SetUpdatedBy(user.idOrSelf(updatedBy...))
	user.SetUpdatedAt()
}

// GenAndSetNameKey generates and sets a NameKey for the model
// FIX: This method must be extracted
// to a generic interface embedded in all models.
func (user *User) GenAndSetNameKey() string {
	nk := userModelName
	if user.Username != "" {
		nk = user.Username
	}
	sfx := GenRandomSufix()
	nk = fmt.Sprintf("%s-%s", nk, sfx)
	nk = strings.ToLower(Trim(nk, maxNameKeyLength))
	user.NameKey = nk
	return user.NameKey
}

// UpdatePasswordDigest creates a password digest from current password.
func (user *User) UpdatePasswordDigest() error {
	if strings.Trim(user.Password, " ") == "" {
		return nil
	}
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(hpass)
	return nil
}

// CalcGeoHash calculates geohash using recever's lat and lng.
func (user *User) CalcGeoHash() (string, error) {
	return "#notimplemented", errors.New("CalcGeoHash not implemented")
}

// CalcAndSetGeoHash calculates and sets geohash using recever's lat and lng.
func (user *User) CalcAndSetGeoHash() (string, error) {
	gh, err := user.CalcGeoHash()
	if err == nil {
		user.Geohash = ""
	}
	return gh, err
}

func (user *User) idOrSelf(userID ...uuid.UUID) uuid.UUID {
	if len(userID) > 0 {
		return userID[0]
	}
	if user.Identification.IsNew() {
		return uuid.Nil
	}
	return user.ID
}
