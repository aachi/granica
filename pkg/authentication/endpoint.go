/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	c "gitlab.com/mikrowezel/backend/granica/internal/config"
)

func makeSignUpEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signUpRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		user, err := svc.SignUp(req.Username, req.Password, req.Email, req.EmailConfirmation, req.TenantID)
		if err != nil {
			return signUpResponse{user, err.Error()}, nil
		}
		return signUpResponse{user, ""}, nil
	}
}

func makeCancelEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		err := svc.Cancel(req.Username, req.Password, req.TenantID)
		if err != nil {
			return cancelResponse{err.Error()}, nil
		}
		return cancelResponse{""}, nil
	}
}

func makeSignInEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signInRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		user, err := svc.SignIn(req.Username, req.Password, req.TenantID)
		if err != nil {
			return signInResponse{user, err.Error()}, nil
		}
		return signInResponse{user, ""}, nil
	}
}

func makeSignOutEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signOutRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		err := svc.SignOut(req.Username, req.TenantID)
		if err != nil {
			return signOutResponse{err.Error()}, nil
		}
		return signOutResponse{""}, nil
	}
}

func makeCreateEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		user, err := svc.Create(req.Username, req.Password, req.Email, req.TenantID)
		if err != nil {
			return createResponse{user, err.Error()}, nil
		}
		return createResponse{user, ""}, nil
	}
}

func makeUpdateEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		err := svc.Update(req.Username, req.Password, req.PasswordConfirmation,
			req.Email, req.EmailConfirmation, req.Description,
			req.GivenName, req.MiddleNames, req.FamilyName,
			req.TenantID)
		if err != nil {
			return updateResponse{err.Error()}, nil
		}
		return updateResponse{""}, nil
	}
}

func makeRemoveEndpoint(svc GranicaService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(removeRequest)
		reqs := fmt.Sprintf("Req: %+v", req)
		svc.Logger().Log("level", c.LogLevel.Debug, "req", reqs)
		err := svc.Remove(req.Username, req.Email, req.TenantID)
		if err != nil {
			return removeResponse{err.Error()}, nil
		}
		return removeResponse{""}, nil
	}
}
