package concrete

import (
	"github.com/Vackhan/metrics/internal/server"
	endpoint "github.com/Vackhan/metrics/internal/server/httpserver/internal/endpoints/update"
	"github.com/Vackhan/metrics/internal/server/storage/memory/update"
)

func WithHTTPServer() server.Server {
	s := &httpServer{}
	s.SetEndpoints(
		endpoint.NewUpdateEndpoint(update.NewMemStorage()),
	)
	return s
}
