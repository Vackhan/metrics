package server

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/pkg/storage/memory/update"
	endpoint "github.com/Vackhan/metrics/internal/server/standardserver/endpoints/update"
)

func WithStandardServer() server.Server {
	s := &httpServer{}
	s.SetEndpoints(
		endpoint.NewUpdateEndpoint(update.NewUpdateMemStorage()),
	)
	return s
}
