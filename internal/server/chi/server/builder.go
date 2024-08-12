package server

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/chi/endpoints/mainpage"
	updendpoint "github.com/Vackhan/metrics/internal/server/chi/endpoints/update"
	"github.com/Vackhan/metrics/internal/server/chi/endpoints/value"
	updmemory "github.com/Vackhan/metrics/internal/server/pkg/storage/memory/update"
)

func WithChiServer() server.Server {
	s := &chiServer{}
	storage := updmemory.NewUpdateMemStorage()
	s.SetEndpoints(
		updendpoint.NewUpdateEndpoint(storage),
		value.NewValueEndpoint(storage),
		mainpage.NewMainpageEndpoint(storage),
	)
	return s
}
