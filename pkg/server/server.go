package server

import (
	_ "clean-gin-template/config"
	"context"
	"net/http"
	"time"
)

const (
	DefaultReadTimeout      = 5 * time.Second
	DefaultWriteTimeout     = 5 * time.Second
	DefaultAddr             = ":80"
	DefaultShutdownTimeout  = 3 * time.Second
	DefaultCertFileLocation = "./config/ssl/server.crt"
	DefaultKeyFileLocation  = "./config/ssl/server.key"
)

// Server initialization
type Server interface {
	run()
	runTLS()
	Notify() <-chan error
	Shutdown() error
}

type server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	scheme          string
	certFileLoc     string
	keyFileLoc      string
}

// New -.
func New(handler http.Handler, opts ...Option) Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		Addr:         DefaultAddr,
	}

	// http is the default server type
	s := &server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: DefaultShutdownTimeout,
		certFileLoc:     DefaultCertFileLocation,
		keyFileLoc:      DefaultKeyFileLocation,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.run()
	//s.runTLS()

	return s
}

func (s *server) run() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *server) runTLS() {
	go func() {
		s.notify <- s.server.ListenAndServeTLS(s.certFileLoc, s.keyFileLoc)
		close(s.notify)
	}()
}

// Notify -.
func (s *server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
