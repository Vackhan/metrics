package main

import (
	"github.com/Vackhan/metrics/internal/server"
	httpConcrete "github.com/Vackhan/metrics/internal/server/httpServer/concrete"
	"log"
)

func main() {
	err := server.NewServer(httpConcrete.WithHttpServer(), ":8080").Run()
	if err != nil {
		log.Println(err)
		return
	}
}
