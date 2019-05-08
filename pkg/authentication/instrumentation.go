package authentication

import (
	"fmt"
	"time"

	// "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	m "gitlab.com/mikrowezel/backend/granica/pkg/models"
)

type instrumentationMiddleware struct {
	logger         log.Logger
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           GranicaService
}

// SignUp is an instrumentation middleware wrapper over another interface implementation of Authenticate.
func (mw instrumentationMiddleware) SignUp(username, password, email, emailConfirmation, tenantID string) (output *m.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SignUp", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.SignUp(username, password, email, emailConfirmation, tenantID)
}

// Cancel is an instrumentation middleware wrapper over another interface implementation of Cancel.
func (mw instrumentationMiddleware) Cancel(username, password, tenantID string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Cancel", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.Cancel(username, password, tenantID)
}

// SignIn is an instrumentation middleware wrapper over another interface implementation of SignIn.
func (mw instrumentationMiddleware) SignIn(username, password, tenantID string) (output *m.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SignIn", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.SignIn(username, password, tenantID)
}

// SignOut is an instrumentation middleware wrapper over another interface implementation of SignOut.
func (mw instrumentationMiddleware) SignOut(username, tenantID string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SignOut", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.SignOut(username, tenantID)
}

// Create is an instrumentation middleware wrapper over another interface implementation of Create.
func (mw instrumentationMiddleware) Create(username, password, email, tenantID string) (output *m.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Create", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.Create(username, password, email, tenantID)
}

// Update is an instrumentation middleware wrapper over another interface implementation of Create.
func (mw instrumentationMiddleware) Update(username, password, passwordConfirmation,
	email, emailConfirmation, description,
	givenName, middleNames, familyName, tenantID string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Update", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.Update(username, password, passwordConfirmation, email,
		emailConfirmation, description, givenName, middleNames, familyName, tenantID)
}

// Remove is an instrumentation middleware wrapper over another interface implementation of Remove.
func (mw instrumentationMiddleware) Remove(username, email, tenantID string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Remove", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.next.Remove(username, email, tenantID)
}

// Remove is an instrumentation middleware wrapper over another interface implementation of Remove.
func (mw instrumentationMiddleware) Logger() log.Logger {
	return mw.logger
}
