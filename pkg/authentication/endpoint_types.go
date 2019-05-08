/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	m "gitlab.com/mikrowezel/backend/granica/pkg/models"
)

// Request & response

// Sign up
type signUpRequest struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	EmailConfirmation string `json:"emailConfirmation"`
	TenantID          string
}

type signUpResponse struct {
	User *m.User `json:"user,omitempty"`
	Err  string  `json:"error,omitempty"`
}

// Cancel
type cancelRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	TenantID string
}

type cancelResponse struct {
	Err string `json:"error,omitempty"`
}

// Sign in
type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	TenantID string
}

type signInResponse struct {
	User *m.User `json:"user,omitempty"`
	Err  string  `json:"error,omitempty"`
}

// Sign out
type signOutRequest struct {
	Username string `json:"username"`
	TenantID string
}

type signOutResponse struct {
	Err string `json:"error,omitempty"`
}

// Create
type createRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	TenantID string
}

type createResponse struct {
	User *m.User `json:"user,omitempty"`
	Err  string  `json:"error,omitempty"`
}

// Update
type updateRequest struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Email                string `json:"email"`
	EmailConfirmation    string `json:"emailConfirmation"`
	Description          string `json:"description"`
	GivenName            string `json:"givenName"`
	MiddleNames          string `json:"middleNames"`
	FamilyName           string `json:"familyName"`
	TenantID             string
}

type updateResponse struct {
	Err string `json:"error,omitempty"`
}

// Remove
type removeRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	TenantID string
}

type removeResponse struct {
	Err string `json:"error,omitempty"`
}
