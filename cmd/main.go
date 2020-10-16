package main

import (
	"log"
	"os"
	"time"
)

type logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type server struct {
	address  string
	insecure bool
	log      logger
	stopCh   chan struct{}
}

func (s *server) run() {
	s.log.Printf("listening on %s", s.address)
	<-s.stopCh
	s.log.Println("done")
}

func main() {
	s := server{
		address:  "0.0.0.0:8080",
		insecure: false,
		log:      log.New(os.Stdout, "[server ]", log.LstdFlags),
		stopCh:   make(chan struct{}),
	}

	go func() {
		time.Sleep(time.Second * 4)
		close(s.stopCh)
	}()

	s.run()
}
