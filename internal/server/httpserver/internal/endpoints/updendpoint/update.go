package updendpoint

import (
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/internal/functionality/upd"
	"github.com/Vackhan/metrics/internal/server/internal/runerr"
	"github.com/Vackhan/metrics/internal/server/internal/storage"
	"net/http"
)

type Update struct {
	repo storage.UpdateRepo
}

func (u *Update) GetURL() string {
	return "/upd/"
}
func (u *Update) GetFunctionality(repos ...server.Repository) any {
	return updateFunc(repos[0])
}

func updateFunc(repo server.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		command := upd.NewUpdate(repo)
		err := command.DoUpdate(r.URL.Path)
		if err != nil {
			switch err {
			case runerr.ErrEmptyMetricName:
				w.WriteHeader(http.StatusNotFound)
			case runerr.ErrWrongMetricType:
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
