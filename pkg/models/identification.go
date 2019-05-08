/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package models

import (
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"time"

	"github.com/google/uuid"
)

const (
	identificationModelName = "identification"
	maxNameKeyLength        = 64
	maxSlugLength           = 128
)

// Identification - Common properties for identification models.
// Id is the system primary key.
// NameKey is a human-meaningful ID.
// Slug is a human-readable identifier mainly used as the last path segment
// in URLs.
type Identification struct {
	ID       uuid.UUID `bson:"_id,omitempty" json:"_id,omitempty"`
	NameKey  string    `bson:"name_key" json:"nameKey"`
	Slug     string    `bson:"slug" json:"slug"`
	TenantID string    `bson:"tenant_id" json:"tenantID"`
}

// func MakeIdentification(nameKey, tenantID string) Identification {
// 	i := Identification{
// 		ID:       null.ZeroUUID(),
// 		NameKey:  makeNameKey(nameKey),
// 		TenantID: null.ToNullsString(tenantID),
// 	}
// 	i.UpdateSlug()
// 	return i
// }

func makeNameKey(nameKey string) string {
	nk := identificationModelName
	if nameKey != "" {
		nk = nameKey
	}
	return nk
}

// UpdateSlug updates the model instance slug.
func (i *Identification) UpdateSlug() {
	u := i.slugSufix()
	nk := i.NameKey
	nk = strings.ToLower(fmt.Sprintf("%s-%s", nk, u))
	nk = strings.ToLower(Trim(nk, maxSlugLength))
	i.Slug = nk
}

func (i *Identification) slugSufix() string {
	id := i.ID.String()
	return b64.StdEncoding.EncodeToString([]byte(id))
}

// GenRandomSufix generates a random sufix for NameKey.
func GenRandomSufix() string {
	return RandomString(12)
}

// RandomString generates a random string of specific length
func RandomString(length int) string {
	var p = "abcdefghijklmnopqrstuvwxyz0123456789$+!*"
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = p[rand.Intn(len(p))]
	}
	return string(bytes)
}

// Trim right-trim a string to the specified length
func Trim(str string, length int) string {
	res := str
	if len(str) > length {
		res = str[0:length]
	}
	return res
}

// SetTenant sets the tenantID.
func (i *Identification) SetTenant(tenantID string) {
	i.TenantID = tenantID
}

// IsNew - True if models has not been persisted yet.
func (i *Identification) IsNew() bool {
	return i.ID == uuid.Nil && i.ID != uuid.Nil
}

// GetID - Get ID.
func (i *Identification) GetID() string {
	return i.ID.String()
}

// SetID - Set ID.
func (i *Identification) SetID(id uuid.UUID) {
	i.ID = id
}

// GenAndSetID - Sets an UUID as ID if its empty.
func (i *Identification) GenAndSetID() {
	i.ID = uuid.New()
}

// GetPK retrieves compund primary key: id + namme key.
func (i *Identification) GetPK() (id uuid.UUID, nameKey string) {
	return i.ID, i.NameKey
}

// IDBytes returns hex decoded of an ID of type UUID as a string representation.
// NOTE: Current implementation of persistence is based on MongoDB,
// But we want to have the chance to also store a vendor independent ID.
// This helper function lets us to construct a finder for an UUID type field.
func (i *Identification) IDBytes() (interface{}, error) {
	hexStr := strings.Replace(i.ID.String(), "-", "", -1)
	ft := hexStr[6:8] + hexStr[4:6] + hexStr[2:4] + hexStr[0:2]
	sd := hexStr[10:12] + hexStr[8:10]
	td := hexStr[14:16] + hexStr[12:14]
	fh := hexStr[16:len(hexStr)]

	hexStr = fmt.Sprintf("%s%s%s%s", ft, sd, td, fh)

	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Match - Custom model comparator.
func (i *Identification) Match(tc Identification) bool {
	r := i.ID == tc.ID &&
		i.Slug == tc.Slug &&
		i.NameKey == tc.NameKey &&
		i.TenantID == tc.TenantID
	return r
}

// MarshalJSON - Custom MarshalJSON function.
func (i *Identification) MarshalJSON() ([]byte, error) {
	type Alias Identification
	return json.Marshal(&struct {
		*Alias
		ID       uuid.UUID `json:"id"`
		Slug     string    `json:"slug"`
		TenantID string    `json:"tenantID"`
	}{
		Alias:    (*Alias)(i),
		ID:       i.ID,
		Slug:     i.Slug,
		TenantID: i.TenantID,
	})
}

// UnmarshalJSON - Custom UnmarshalJSON function.
func (i *Identification) UnmarshalJSON(data []byte) error {
	type Alias Identification
	aux := &struct {
		*Alias
		ID       uuid.UUID `json:"id"`
		Slug     string    `json:"slug"`
		TenantID string    `json:"tenantID"`
	}{
		Alias:    (*Alias)(i),
		ID:       i.ID,
		Slug:     i.Slug,
		TenantID: i.TenantID,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
