package server

import (
	"errors"
	"net/http"
)

var WrongHandlerType = errors.New("wrong handler type")

type Server interface {
	Run() error
	SetUrlListener(url string)
	SetEndpoints(endPointList ...Endpoint)
}

type Handler func(w http.ResponseWriter, r *http.Request)
type Endpoint interface {
	GetUrl() string
	GetFunctionality(repos ...Repository) any
}

type Repository interface {
	ImplementRepository()
}
