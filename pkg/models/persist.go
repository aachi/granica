/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package models

// Persist is a Generic interface for persistible models
type Persist struct{}

// SetCreateValues sets creation values
func (p *Persist) SetCreateValues(tenantID string) {
	// Interface method
}

// SetUpdateValues sets creation values
func (p *Persist) SetUpdateValues() {
	// Interface method
}

// GenNameKey generates a NameKey for the model
func (p *Persist) GenNameKey() string {
	return ""
}
