package updendpoint

import (
	"github.com/Vackhan/metrics/internal/server/internal/functionality/upd"
	"github.com/Vackhan/metrics/internal/server/internal/runerr"
	"github.com/Vackhan/metrics/internal/server/internal/storage"
	"net/http"
)

type Update struct {
	repo storage.UpdateRepo
}

func (u *Update) GetURL() string {
	return "/update/"
}
func (u *Update) GetFunctionality() any {
	return updateFunc(u.repo)
}

func updateFunc(repo storage.UpdateRepo) func(w http.ResponseWriter, r *http.Request) {
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
		_, err = w.Write(nil)
		if err != nil {
			return
		}
	}
}

func NewUpdateEndpoint(repo storage.UpdateRepo) *Update {
	return &Update{repo}
}
