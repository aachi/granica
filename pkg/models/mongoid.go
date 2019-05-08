/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoID - Common properties for Mongo persisted models.
type MongoID struct {
	ObjID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

// IsNew - True if models has not been persisted yet.
func (m *MongoID) IsNew() bool {
	return m.ObjID == primitive.NilObjectID
}

// Match - Custom model comparator.
func (m *MongoID) Match(tc MongoID) bool {
	r := m.ObjID == tc.ObjID
	return r
}

// MarshalJSON - Custom MarshalJSON function.
func (m *MongoID) MarshalJSON() ([]byte, error) {
	type Alias MongoID
	return json.Marshal(&struct {
		*Alias
		ObjID primitive.ObjectID `json:"objID"`
	}{
		Alias: (*Alias)(m),
		ObjID: m.ObjID,
	})
}

// UnmarshalJSON - Custom UnmarshalJSON function.
func (m *MongoID) UnmarshalJSON(data []byte) error {
	type Alias MongoID
	aux := &struct {
		*Alias
		ObjID primitive.ObjectID `json:"objID"`
	}{
		Alias: (*Alias)(m),
		ObjID: m.ObjID,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
