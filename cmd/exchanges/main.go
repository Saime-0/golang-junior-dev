package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/app"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.Info("start")
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("config loaded")
	level, err := logrus.ParseLevel(cfg.LogLevel())
	if err != nil {
		logrus.Fatalf("main - parse log level: %v", err)
	}
	logrus.SetLevel(level)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err = app.Start(ctx, cfg); err != nil {
			logrus.Fatal(err)
		}
		os.Exit(0)
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	logrus.Info("main - signal: " + (<-interrupt).String())
	cancel()
	time.Sleep(3 * time.Second)
}
