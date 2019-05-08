/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	c "gitlab.com/mikrowezel/backend/granica/internal/config"
	m "gitlab.com/mikrowezel/backend/granica/pkg/models"
)

type loggingMiddleware struct {
	logger log.Logger
	next   GranicaService
}

// SignUp is a logging middleware wrapper over another interface implementation of SignUp.
func (mw loggingMiddleware) SignUp(username, password, email, emailConfirmation, tenantID string) (output *m.User, err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s, %s, %s, %s}", username, password, email, emailConfirmation, tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "SignUp",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.SignUp(username, password, email, emailConfirmation, tenantID)
	return
}

// Cancel is a logging middleware wrapper over another interface implementation of Cancel.
func (mw loggingMiddleware) Cancel(username, password, tenantID string) (err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s, %s}", username, "********", tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "Cancel",
			"input", input,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Cancel(username, password, tenantID)
}

// SignIn is a logging middleware wrapper over another interface implementation of SingnIn.
func (mw loggingMiddleware) SignIn(username, password, tenantID string) (output *m.User, err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s, %s}", username, password, tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "SignIn",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.SignIn(username, password, tenantID)
	return
}

// SignOut is a logging middleware wrapper over another interface implementation of SingnOut.
func (mw loggingMiddleware) SignOut(username, tenantID string) (err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s}", username, tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "SignOut",
			"input", input,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.SignOut(username, tenantID)
}

// Create is a logging middleware wrapper over another interface implementation of Create.
func (mw loggingMiddleware) Create(username, password, email, tenantID string) (output *m.User, err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s, %s, %s}", username, password, email, tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "Create",
			"input", input,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Create(username, password, email, tenantID)
	return
}

// Update is a logging middleware wrapper over another interface implementation of Create.
func (mw loggingMiddleware) Update(username, password, passwordConfirmation,
	email, emailConfirmation, description, givenName, middleNames, familyName,
	tenantID string) (err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s, %s, %s}", username, password, email, tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "Update",
			"input", input,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.Update(username, password, passwordConfirmation, email,
		emailConfirmation, description, givenName, middleNames, familyName, tenantID)
	return
}

// Remove is a logging middleware wrapper over another interface implementation of Remove.
func (mw loggingMiddleware) Remove(username, email, tenantID string) (err error) {
	defer func(begin time.Time) {
		input := fmt.Sprintf("{%s, %s, %s}", username, email, tenantID)
		mw.logger.Log(
			"level", c.LogLevel.Info,
			"method", "Remove",
			"input", input,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Remove(username, email, tenantID)
}

// Remove is a logging middleware wrapper over another interface implementation of Remove.
func (mw loggingMiddleware) Logger() log.Logger {
	return mw.logger
}
