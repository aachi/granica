/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Audit - Common properties for auditable models.
type Audit struct {
	CreatedBy uuid.UUID `bson:"created_by" json:"createdBy"`
	UpdatedBy uuid.UUID `bson:"updated_by" json:"updatedBy"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt"`
}

// SetCreatedBy set the CreatedAt date.
func (a *Audit) SetCreatedBy(id uuid.UUID) {
	a.CreatedBy = id
}

// SetUpdatedBy set the UpdatedAt date.
func (a *Audit) SetUpdatedBy(id uuid.UUID) {
	a.UpdatedBy = id
}

// SetCreatedAt set the CreatedAt date.
func (a *Audit) SetCreatedAt() {
	a.CreatedAt = time.Now()
}

// SetUpdatedAt set the UpdatedAt date.
func (a *Audit) SetUpdatedAt() {
	a.UpdatedAt = time.Now()
}

// Match - Custom model comparator.
func (a *Audit) Match(tc Audit) bool {
	r := a.CreatedBy == tc.CreatedBy &&
		a.UpdatedBy == tc.UpdatedBy &&
		a.CreatedAt == tc.CreatedAt &&
		a.UpdatedAt == tc.UpdatedAt
	return r
}

// MarshalJSON - Custom MarshalJSON function.
func (a *Audit) MarshalJSON() ([]byte, error) {
	type Alias Audit
	return json.Marshal(&struct {
		*Alias
		CreatedBy string `json:"createdBy"`
		UpdatedBy string `json:"updatedBy"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
	}{
		Alias:     (*Alias)(a),
		CreatedBy: a.CreatedBy.String(),
		UpdatedBy: a.UpdatedBy.String(),
		CreatedAt: a.CreatedAt.Unix(),
		UpdatedAt: a.UpdatedAt.Unix(),
	})
}

// UnmarshalJSON - Custom UnmarshalJSON function.
func (a *Audit) UnmarshalJSON(data []byte) error {
	type Alias Audit
	aux := &struct {
		*Alias
		CreatedBy string `json:"createdBy"`
		UpdatedBy string `json:"updatedBy"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
	}{
		Alias:     (*Alias)(a),
		CreatedBy: a.CreatedBy.String(),
		UpdatedBy: a.UpdatedBy.String(),
		CreatedAt: a.CreatedAt.Unix(),
		UpdatedAt: a.UpdatedAt.Unix(),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	tc := time.Unix(aux.CreatedAt, 0)
	tu := time.Unix(aux.UpdatedAt, 0)
	a.CreatedAt = tc
	a.UpdatedAt = tu
	return nil
}
