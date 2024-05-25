package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/httpserver"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/usecase"
	"sync"
)

func Start(ctx context.Context, cfg *config.Config) error {
	var wg sync.WaitGroup

	calcExchangeUsecase := &usecase.CalculatingExchange{}

	httpServer := httpserver.New(
		ctx,
		httpserver.Usecases{
			CalculatingExchange: calcExchangeUsecase,
		},
		cfg.Host(),
		cfg.Port(),
	)
	wg.Add(1)
	go func() {
		<-httpServer.Done()
		wg.Done()
	}()
	errChan := make(chan error)
	go func() {
		var err error
		select { // read notifications from components:
		case err = <-httpServer.Notify():
		}
		select { // send to err chan
		case errChan <- fmt.Errorf("app - Run - httpServer.Notify: %w", err):
		default: // <- non blocking
		}
		close(errChan)
	}()
	select {
	case <-ctx.Done():
		logrus.Info("app - Run - receive ctx.Done, wait when components stop the work")
		wg.Wait()
		logrus.Info("app - Run - components done the work")
		return nil
	case err := <-errChan:
		return err
	}
}
