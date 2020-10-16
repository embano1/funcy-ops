package server

type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Server struct {
	address  string
	insecure bool
	log      Logger
	stopCh   chan struct{}
}

func New(address string, insecure bool, log Logger) *Server {
	// apply defaults, check for empty values, etc.
	return &Server{address: address, insecure: insecure, log: log, stopCh: make(chan struct{})}
}

func (s *Server) Run() {
	s.log.Printf("listening on %s", s.address)
	<-s.stopCh
	s.log.Println("done")
}

func (s *Server) Stop() {
	close(s.stopCh)
}
