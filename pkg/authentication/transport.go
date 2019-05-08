/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	"context"
	"encoding/json"

	"net/http"
	"net/url"

	httptransport "github.com/go-kit/kit/transport/http"
	"gitlab.com/mikrowezel/backend/granica/internal/config"
	c "gitlab.com/mikrowezel/backend/granica/internal/config"
)

const (
	loggingOn         = true
	instrumentationOn = true
	transactionOn     = true
)

// Transport: JSON over HTTP.

// Run - starts the service.
func Run() {

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	go checkSigTerm(cancel)

	// Config
	cfg, err := config.Load()
	checkError(err)

	// Logger
	logger := makeLogger()

	// Service
	svc, err := makeService(ctx, cfg, logger).Init()
	checkError(err)

	http.Handle("/sign-up", SignUpHandler(svc))
	http.Handle("/cancel", CancelHandler(svc))
	http.Handle("/sign-in", SignInHandler(svc))
	http.Handle("/sign-out", SignOutHandler(svc))
	http.Handle("/create", CreateHandler(svc))
	http.Handle("/update", UpdateHandler(svc))
	http.Handle("/remove", RemoveHandler(svc))

	err = http.ListenAndServe(":8080", nil)

	logger.Log("level", c.LogLevel.Error, "msg", err.Error())
}

// SignUpHandler manages signing up process.
func SignUpHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeSignUpEndpoint(svc),
		decodeSignUpRequest,
		encodeResponse,
	)
}

// CancelHandler manages cancel process.
func CancelHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeCancelEndpoint(svc),
		decodeCancelRequest,
		encodeResponse,
	)
}

// SignInHandler manages signing in process.
func SignInHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeSignInEndpoint(svc),
		decodeSignInRequest,
		encodeResponse,
	)
}

// SignOutHandler manages signing out process.
func SignOutHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeSignOutEndpoint(svc),
		decodeSignOutRequest,
		encodeResponse,
	)
}

// CreateHandler manages create user process
func CreateHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)
}

// UpdateHandler manages user update process
func UpdateHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeUpdateEndpoint(svc),
		decodeUpdateRequest,
		encodeResponse,
	)
}

// RemoveHandler manages user removing process
func RemoveHandler(svc GranicaService) *httptransport.Server {
	return httptransport.NewServer(
		makeRemoveEndpoint(svc),
		decodeRemoveRequest,
		encodeResponse,
	)
}

// Decoders
func decodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func decodeCancelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request cancelRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func decodeSignInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request signInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func decodeSignOutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request signOutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request updateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func decodeRemoveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request authenticateRequest = authenticateRequest{}
	var request removeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.TenantID = getTenant(r)
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func getTenant(r *http.Request) string {
	h, err := url.ParseRequestURI("https://" + r.Host)
	if err != nil {
		return ""
	}
	return h.Hostname()
}
