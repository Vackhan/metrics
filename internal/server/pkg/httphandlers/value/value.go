package value

import (
	"fmt"
	"github.com/Vackhan/metrics/internal/server/pkg/functionality/value"
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"net/http"
	"strings"
)

// New обработчик для эндпоинта value
func New(repo storage.UpdateRepo, params func(r *http.Request) (string, string)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		valueFunctional := value.NewValue(repo)
		if params == nil {
			params = standardParamsConsuming
		}
		metricType, metricName := params(r)
		val, err := valueFunctional.GetValueByTypeAndName(metricType, metricName)
		if err != nil {
			switch err {
			case runerr.ErrMetricNotFound:
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write(nil)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%v", val)))
		}
	}
}

func standardParamsConsuming(r *http.Request) (string, string) {
	urlData := strings.Split(r.URL.Path[1:], "/")
	newData := make([]string, 3)
	copy(newData, urlData)
	return newData[1], newData[2]
}
