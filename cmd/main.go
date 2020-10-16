package main

import (
	"log"
	"os"
	"time"

	server "embano1/funcy-ops"
)

func main() {
	logger := log.New(os.Stdout, "[my server] ", log.LstdFlags|log.Lshortfile)
	s, err := server.New()
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	s.Run()
}
