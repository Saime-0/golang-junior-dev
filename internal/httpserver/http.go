package httpserver

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/usecase"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	server *http.Server
	notify chan error
	done   chan struct{}
	uc     Usecases
}

func (s *Server) start(ctx context.Context) {
	go func() {
		<-ctx.Done()
		// send to notify not required, because app.Start
		// handle ctx.Done earlier
		err := s.server.Shutdown(context.Background())
		if err != nil {
			logrus.Printf("httpserver - Server - closed with error: %v", err)
		}
	}()
	err := s.server.ListenAndServe()
	// ignore ErrServerClose handling, because in
	// this case - app.Start handle ctx.Done earlier
	s.notify <- fmt.Errorf("httpserver - Server - server.ListenAndServe: %w", err)
	close(s.notify)
}

type Usecases struct {
	CalculatingExchange *usecase.CalculatingExchange
}

func New(ctx context.Context, usecases Usecases, host string, port int) *Server {
	mux := http.NewServeMux()

	s := &Server{
		server: &http.Server{
			Addr:        host + ":" + strconv.Itoa(port),
			Handler:     mux,
			ReadTimeout: 30 * time.Second,
		},
		uc:     usecases,
		notify: make(chan error),
		done:   make(chan struct{}),
	}

	mux.HandleFunc("POST /calculating-exchange", s.calculatingExchange())

	s.start(ctx)
	return s
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Done() <-chan struct{} {
	return s.done
}
