package main

import (
	"github.com/Vackhan/metrics/internal/server"
	chiServer "github.com/Vackhan/metrics/internal/server/chi/server"
	"log"
)

func main() {
	err := server.NewServer(
		chiServer.WithChiServer(),
		":8080",
	).Run()
	if err != nil {
		log.Println(err)
		return
	}
}
