package server

import (
	"net/http"
)

type Server interface {
	Run() error
	SetUrlListener(url string)
	SetEndpoints(endPointList ...Endpoint)
}

type Handler func(w http.ResponseWriter, r *http.Request)
type Endpoint interface {
	GetURL() string
	GetFunctionality(repos ...Repository) any
}

type Repository interface {
	ImplementRepository()
}
