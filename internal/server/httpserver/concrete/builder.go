package concrete

import (
	"github.com/Vackhan/metrics/internal/server"
	endpoint "github.com/Vackhan/metrics/internal/server/httpserver/endpoints/update"
	"github.com/Vackhan/metrics/internal/server/pkg/storage/memory/update"
)

func WithHTTPServer() server.Server {
	s := &httpServer{}
	s.SetEndpoints(
		endpoint.NewUpdateEndpoint(update.NewUpdateMemStorage()),
	)
	return s
}
