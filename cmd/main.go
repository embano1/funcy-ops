package main

import (
	"log"
	"os"
	"time"

	server "embano1/funcy-ops"
)

func main() {
	logger := log.New(os.Stdout, "[server] ", log.LstdFlags|log.Lshortfile)
	cfg := server.Config{
		Address:  "0.0.0.0:8080",
		Insecure: false,
		Log:      logger,
		StopCh:   make(chan struct{}),
	}

	s, err := server.New(&cfg)
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	s.Run()
}
