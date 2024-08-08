package update

import (
	"github.com/Vackhan/metrics/internal/server/functionality/update"
	"github.com/Vackhan/metrics/internal/server/runerr"
	"github.com/Vackhan/metrics/internal/server/storage"
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
		command := update.NewUpdate(repo)
		err := command.DoUpdate(r.URL.Path)
		if err != nil {
			switch err {
			case runerr.ErrEmptyMetricName:
				w.WriteHeader(http.StatusNotFound)
			case runerr.ErrWrongMetricType:
				w.WriteHeader(http.StatusBadRequest)
			case runerr.ErrWrongURL:
				w.WriteHeader(http.StatusNotFound)
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
