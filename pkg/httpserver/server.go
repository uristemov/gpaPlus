package httpserver

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

type server struct {
	logger          *zap.SugaredLogger
	shutdownTimeout time.Duration
	client          *http.Server
	listener        net.Listener
	isReady         bool
}

func NewServer(port int, shutdownTimeout time.Duration, logger *zap.SugaredLogger, handler http.Handler) (*server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("connot bind HTTP server '%d': %v", port, err)
	}

	return &server{
		client: &http.Server{
			Handler: handler,
		},
		listener:        listener,
		shutdownTimeout: shutdownTimeout,
		logger:          logger,
		isReady:         false,
	}, nil
}

func (s *server) Ready() error {
	if s.isReady {
		return nil
	}

	return errors.New("[INFO] Server is not ready!")
}

func (s *server) Stop() error {
	s.isReady = false
	s.logger.Infof("[%s] HTTP server is stopping...", s.listener.Addr().String())

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	s.client.SetKeepAlivesEnabled(false)

	if err := s.client.Shutdown(ctx); err != nil {
		return fmt.Errorf("cannot stop HTTP server: %w", err)
	}

	s.logger.Infof("[%s] HTTP server was stopped", s.listener.Addr().String())

	return nil
}

func (s *server) Run() {

	go func() {
		s.isReady = true
		//s.logger.Infof("[%s] HTTP server is run", s.listener.Addr().String())

		if err := s.client.Serve(s.listener); err != nil {
			s.isReady = false
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			s.logger.Errorf("[%s] HTTP server was stopped with error: %s", s.listener.Addr().String(), err)
		}
	}()
}
