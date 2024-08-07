package update_endpoint

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/internal/functionality/update_endpoint"
	"github.com/Vackhan/metrics/internal/server/internal/functionalityErrors"
	"github.com/Vackhan/metrics/internal/server/internal/storage"
	"net/http"
)

type Update struct {
	repo storage.UpdateRepo
}

func (u *Update) GetUrl() string {
	return "/update/"
}
func (u *Update) GetFunctionality(repos ...server.Repository) any {
	return updateFunc(repos[0])
}

func updateFunc(repo server.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		command := update_endpoint.NewUpdate(repo)
		err := command.DoUpdate(r.URL.Path)
		if err != nil {
			switch err {
			case functionalityErrors.ErrEmptyMetricName:
				w.WriteHeader(http.StatusNotFound)
			case functionalityErrors.ErrWrongMetricType:
				w.WriteHeader(http.StatusBadRequest)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Write(nil)
	}
}

func NewUpdateEndpoint(repo storage.UpdateRepo) *Update {
	return &Update{repo}
}
