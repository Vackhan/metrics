package server

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type chiServer struct {
	url       string
	endpoints []server.Endpoint
}

func (h *chiServer) SetURLListener(url string) {
	h.url = url
}

func (h *chiServer) SetEndpoints(e ...server.Endpoint) {
	h.endpoints = e
}

func (h *chiServer) Run() error {
	r := chi.NewRouter()
	log.Println(h.endpoints)
	for _, e := range h.endpoints {
		f := e.GetFunctionality()
		switch v := f.(type) {
		case func(r chi.Router):
			r.Route(e.GetURL(), v)
		default:
			return runerr.ErrWrongHandlerType
		}
		log.Println("listen to " + e.GetURL())
	}
	log.Println("run server")
	return http.ListenAndServe(h.url, r)
}
