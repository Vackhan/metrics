package main

import (
	"github.com/Vackhan/metrics/internal/server"
	chiServer "github.com/Vackhan/metrics/internal/server/chiserver/server"
	"github.com/Vackhan/metrics/internal/server/pkg/flags"
	"github.com/spf13/pflag"
	"log"
)

func main() {
	addressAndPort, err := flags.GetAddress()
	if err != nil {
		log.Fatalln(err)
	}
	pflag.Parse()
	err = server.NewServer(
		chiServer.WithChiServer(),
		addressAndPort.String(),
	).Run()
	if err != nil {
		log.Println(err)
		return
	}
}
