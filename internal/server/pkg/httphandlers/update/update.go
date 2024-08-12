package update

import (
	"github.com/Vackhan/metrics/internal/server/pkg/functionality/update"
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"net/http"
)

// New обработчик для update серверов, совместимых со стандартным сервером go
func New(repo storage.UpdateRepo) func(w http.ResponseWriter, r *http.Request) {
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
