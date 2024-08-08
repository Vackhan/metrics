package concrete

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/httpserver/internal/endpoints/updendpoint"
	"github.com/Vackhan/metrics/internal/server/internal/storage/memory/update"
)

func WithHTTPServer() server.Server {
	s := &httpServer{}
	s.SetEndpoints(
		updendpoint.NewUpdateEndpoint(update.NewMemStorage()),
	)
	return s
}
