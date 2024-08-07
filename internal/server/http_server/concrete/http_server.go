package concrete

import (
	"github.com/Vackhan/metrics/internal/server"
	"net/http"
)

type httpServer struct {
	url       string
	endpoints []server.Endpoint
}

func (h *httpServer) SetUrlListener(url string) {
	h.url = url
}

func (h *httpServer) SetEndpoints(e ...server.Endpoint) {
	h.endpoints = e
}

func (h *httpServer) Run() error {
	mux := http.NewServeMux()
	for _, e := range h.endpoints {
		f, ok := e.GetFunctionality().(httpServerHandler)
		if !ok {
			return server.WrongHandlerType
		}
		mux.HandleFunc(e.GetUrl(), f)
	}
	return http.ListenAndServe(h.url, mux)
}

type httpServerHandler func(w http.ResponseWriter, r *http.Request)
