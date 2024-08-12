package value

import (
	"github.com/Vackhan/metrics/internal/server/pkg/httphandlers/value"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Value struct {
	repo storage.UpdateRepo
}

func (u *Value) GetURL() string {
	return "/value"
}
func (u *Value) GetFunctionality() any {
	return func(r chi.Router) {
		r.Get(
			"/{type}/{name}",
			value.New(
				u.repo,
				func(r *http.Request) (string, string) {
					return chi.URLParam(r, "type"), chi.URLParam(r, "name")
				},
			),
		)
	}
}

func NewValueEndpoint(repo storage.UpdateRepo) *Value {
	return &Value{repo}
}
