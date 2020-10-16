package server

import (
	"errors"
)

type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Server struct {
	cfg *Config
}

type Config struct {
	Address  string
	Insecure bool
	Log      Logger
	StopCh   chan struct{}
}

func New(cfg *Config) (*Server, error) {
	if cfg == nil {
		return nil, errors.New("empty config")
	}

	// apply defaults, check for empty values, etc.
	if cfg.Insecure {
		return nil, errors.New("security alert")
	}

	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Run() {
	s.cfg.Log.Printf("listening on %s", s.cfg.Address)
	<-s.cfg.StopCh
	s.cfg.Log.Println("done")
}

func (s *Server) Stop() {
	close(s.cfg.StopCh)
}
