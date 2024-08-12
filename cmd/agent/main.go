package main

import (
	"context"
	"fmt"
	"github.com/Vackhan/metrics/internal/agent"
	"github.com/Vackhan/metrics/internal/server/pkg/flags"
	"github.com/spf13/pflag"
	"time"
)

func main() {
	serverAddrAndPort := flags.NewAddress()
	pflag.VarP(serverAddrAndPort, "a", "a", "host and port of the listener")
	ttl := pflag.IntP("ttl", "t", 0, "time to live for agent")
	r := pflag.IntP("r", "r", 10, "server send frequency in seconds")
	p := pflag.IntP("p", "p", 2, "consuming data frequency in seconds")
	pflag.Parse()
	var ctx context.Context
	var cancel context.CancelFunc
	if *ttl > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(*ttl)*time.Second)
		defer cancel()
	} else {
		ctx = context.Background()
	}
	agent.New(
		fmt.Sprintf("http://%s", serverAddrAndPort.String()),
		ctx,
		*r,
		*p,
	).Run()
}
