package mainpage

import (
	handler "github.com/Vackhan/metrics/internal/server/pkg/httphandlers/mainpage"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"github.com/go-chi/chi/v5"
)

type MainPaige struct {
	repo storage.UpdateRepo
}

func (u *MainPaige) GetURL() string {
	return "/"
}
func (u *MainPaige) GetFunctionality() any {
	return func(r chi.Router) {
		r.Get("/", handler.New(u.repo))
	}
}

func NewMainpageEndpoint(repo storage.UpdateRepo) *MainPaige {
	return &MainPaige{repo}
}
