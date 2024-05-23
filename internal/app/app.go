package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/httpserver"
	"sync"
)

func Start(ctx context.Context, cfg *config.Config) error {
	var wg sync.WaitGroup

	httpServer := httpserver.New(ctx, cfg.Host(), cfg.Port())
	wg.Add(1)
	go func() {
		<-httpServer.Done()
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		logrus.Info("app - Run - receive ctx.Done, wait when components stop the work")
		wg.Wait()
		logrus.Info("app - Run - components done the work")
		return nil
	case err := <-httpServer.Notify():
		return fmt.Errorf("app - Run - httpServer.Notify: %w", err)
	}
}
