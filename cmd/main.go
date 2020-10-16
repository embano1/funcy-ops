package main

import (
	"log"
	"os"
	"time"

	server "embano1/funcy-ops"
)

func main() {
	s := server.New("0.0.0.0:8080", false, log.New(os.Stdout, "[server] ", log.LstdFlags|log.Lshortfile))

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	s.Run()
}
