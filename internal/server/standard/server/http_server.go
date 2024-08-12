package server

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"log"
	"net/http"
)

type httpServer struct {
	url       string
	endpoints []server.Endpoint
}

func (h *httpServer) SetURLListener(url string) {
	h.url = url
}

func (h *httpServer) SetEndpoints(e ...server.Endpoint) {
	h.endpoints = e
}

func (h *httpServer) Run() error {
	mux := http.NewServeMux()
	log.Println(h.endpoints)
	for _, e := range h.endpoints {
		f, ok := e.GetFunctionality().(func(w http.ResponseWriter, r *http.Request))
		if !ok {
			return runerr.ErrWrongHandlerType
		}
		mux.HandleFunc(e.GetURL(), f)
		log.Println("listen to " + e.GetURL())
	}
	log.Println("run server")
	return http.ListenAndServe(h.url, mux)
}
