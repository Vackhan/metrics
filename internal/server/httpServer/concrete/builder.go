package concrete

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/httpServer/internal/endpoints/update_endpoint"
	"github.com/Vackhan/metrics/internal/server/internal/storage/memory/update"
)

func WithHttpServer() server.Server {
	s := &httpServer{}
	s.SetEndpoints(
		update_endpoint.NewUpdateEndpoint(update.NewMemStorage()),
	)
	return &httpServer{}
}
