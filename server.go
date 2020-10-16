package server

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	defaultAddress  = "0.0.0.0:8080"
	skipCertWarning = false
	logPrefix       = "[server] "
)

type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Option changes default Server parameters
type Option func(*Server)

// WithInsecure skips certificate verification
func WithInsecure(in bool) Option {
	return func(s *Server) {
		s.insecure = in
	}
}

// WithAddress sets the listener address
func WithAddress(addr string) Option {
	return func(s *Server) {
		s.address = addr
	}
}

// WithLogger sets a custom logger
func WithLogger(log Logger) Option {
	return func(s *Server) {
		s.log = log
	}
}

// Server runs a server
type Server struct {
	address  string
	insecure bool
	log      Logger
	stopCh   chan struct{}
}

// New creates a new server listening on defaultAddress
func New(ops ...Option) (*Server, error) {
	prefix := fmt.Sprintf("%s", logPrefix)
	s := Server{
		address:  defaultAddress,
		insecure: skipCertWarning,
		log:      log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile),
		stopCh:   make(chan struct{}),
	}

	for _, o := range ops {
		o(&s)
	}

	// validate defaults, options, etc.
	if s.insecure {
		return nil, errors.New("security alert")
	}

	return &s, nil
}

// Run starts the server
func (s *Server) Run() {
	s.log.Printf("listening on %s", s.address)
	<-s.stopCh
	s.log.Println("done")
}

// Stop gracefully stops the server
func (s *Server) Stop() {
	close(s.stopCh)
}
