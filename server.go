package server

import (
	"errors"
	"log"
	"os"
	"strings"
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

// ServerOption changes default Server parameters
type ServerOption func(*Server) error

// WithInsecure skips certificate verification
func WithInsecure(in bool) ServerOption {
	return func(s *Server) error {
		if in {
			return errors.New("security alert")
		}
		s.insecure = in
		return nil
	}
}

// WithAddress sets the listener address
func WithAddress(addr string) ServerOption {
	return func(s *Server) error {
		if !strings.Contains(addr, ":") {
			return errors.New("port must be specified")
		}
		s.address = addr
		return nil
	}
}

// WithLogger sets a custom logger
func WithLogger(log Logger) ServerOption {
	return func(s *Server) error {
		if log == nil {
			return errors.New("logger is nil")
		}
		s.log = log
		return nil
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
func New(ops ...ServerOption) (*Server, error) {
	s := Server{
		address:  defaultAddress,
		insecure: skipCertWarning,
		log:      log.New(os.Stdout, logPrefix, log.LstdFlags|log.Lshortfile),
		stopCh:   make(chan struct{}),
	}

	for _, o := range ops {
		err := o(&s)
		if err != nil {
			return nil, err
		}
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
