package main

import "github.com/Vackhan/metrics/internal/agent"

func main() {
	agent.New().Run("localhost:8080")
}
