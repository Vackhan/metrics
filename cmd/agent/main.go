package main

import (
	"context"
	"github.com/Vackhan/metrics/internal/agent"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	agent.New().Run("localhost:8080", ctx)
}
