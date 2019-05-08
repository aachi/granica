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
	l "log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"gitlab.com/mikrowezel/backend/granica/internal/config"
	"gitlab.com/mikrowezel/backend/granica/pkg/authentication/repo"
)

// checkSigTerm listens to sigterm events.
func checkSigTerm(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	l.Printf("[ERROR] service interrupted.")
	cancel()
}

func makeService(ctx context.Context, cfg *config.Config, log log.Logger) *granicaService {
	return &granicaService{
		name:   "Granica",
		ctx:    ctx,
		cfg:    cfg,
		logger: log,
	}
}

// Init a service instance.
func (svc *granicaService) Init() (GranicaService, error) {
	var gs GranicaService

	// Repo
	rerr := make(chan error)
	rrepo := make(chan repo.UserRepo)
	go func() {
		c, err := repo.NewRepo(svc.ctx, svc.cfg, svc.Logger())
		if err != nil {
			rerr <- err
			return
		}
		rerr <- nil
		rrepo <- c
	}()

	if <-rerr != nil {
		return gs, fmt.Errorf("cannot initialize '%s' service", svc.name)
	}

	svc.repo = <-rrepo

	// Middleware
	gs = addLogging(svc, svc.logger)
	gs = addInstrumentation(svc)

	return gs, nil
}

// NewServiceForTests returns a configured sertvice
// Mainly used for tests.
func NewServiceForTests(cfg *config.Config) GranicaService {
	// Context
	ctx, cancel := context.WithCancel(context.Background())
	go checkSigTerm(cancel)

	// Logger
	logger := makeLogger()

	// Config
	cfg, err := config.Load()
	checkError(err)

	// Service
	svc := makeService(ctx, cfg, logger)
	return svc
}

// Middleware
func addLogging(svc GranicaService, logger log.Logger) GranicaService {
	if loggingOn {
		return loggingMiddleware{logger, svc}
	}
	return svc
}

func addInstrumentation(svc GranicaService) GranicaService {
	if instrumentationOn {
		m := instrumentationMeters()
		return instrumentationMiddleware{svc.Logger(), m.ReqCount, m.ReqLatency, m.CountResult, svc}
	}
	return svc
}

func makeLogger() log.Logger {
	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)
	logger.Log("level", config.LogLevel.Info, "message", "Config Logger started.")
	return logger
}

// Start the service.
func (svc *granicaService) Start() {
	go svc.checkCancel()
	// s.repo.Start()
}

func (svc *granicaService) checkCancel() {
	<-svc.ctx.Done()
	// s.repo.Stop()
}
