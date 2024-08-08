package main

import (
	"github.com/Vackhan/metrics/internal/server"
	httpConcrete "github.com/Vackhan/metrics/internal/server/httpserver/concrete"
	"log"
)

func main() {
	err := server.NewServer(
		httpConcrete.WithHTTPServer(),
		":8080",
	).Run()
	if err != nil {
		log.Println(err)
		return
	}
}
