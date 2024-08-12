package main

import (
	"context"
	"fmt"
	"github.com/Vackhan/metrics/internal/agent"
	"github.com/Vackhan/metrics/internal/server/pkg/flags"
	"github.com/spf13/pflag"
	"log"
	"time"
)

func main() {
	serverAddrAndPort, err := flags.GetAddress()
	if err != nil {
		log.Fatalln(err)
	}
	ttlObj, err := flags.GetTTL()
	if err != nil {
		log.Fatalln(err)
	}
	reportInterval, err := flags.GetReportInterval()
	if err != nil {
		log.Fatalln(err)
	}
	pollInterval, err := flags.GetPollInterval()
	if err != nil {
		log.Fatalln(err)
	}
	pflag.Parse()
	var ctx context.Context
	var cancel context.CancelFunc
	if ttlObj.Duration > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), ttlObj.Duration*time.Second)
		defer cancel()
	} else {
		ctx = context.Background()
	}
	agent.New(
		fmt.Sprintf("http://%s", serverAddrAndPort.String()),
		ctx,
		reportInterval.Interval,
		pollInterval.Interval,
	).Run()
}
