package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/saime-0/cute-chat-backend/internal/config"
	"github.com/saime-0/cute-chat-backend/internal/httpserver"
	"github.com/saime-0/cute-chat-backend/internal/repository/postgres"
	"log"
	"sync"
)

func Start(ctx context.Context, cfg *config.Config) error {
	var wg sync.WaitGroup
	db, err := pgx.Connect(ctx, cfg.DbConnString())
	if err != nil {
		return fmt.Errorf("app - Start - pgx.Connect: %w", err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		db.Close(context.Background())
	}()

	commonRepository := postgres.NewCommonRepository(db)

	httpServer := httpserver.New(ctx, cfg.Listen())
	wg.Add(1)
	go func() {
		<-httpServer.Done()
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		log.Println("app - Run - receive ctx.Done, wait when components stop the work")
		wg.Wait()
		log.Println("app - Run - components done the work")
		return nil
	case err = <-httpServer.Notify():
		return fmt.Errorf("app - Run - httpServer.Notify: %w", err)
	}
}
