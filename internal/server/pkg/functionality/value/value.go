package value

import (
	"github.com/Vackhan/metrics/internal/server/pkg/functionality"
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"log"
)

type Value struct {
	storage storage.UpdateRepo
}

func (v Value) GetValueByTypeAndName(metricType, metricName string) (any, error) {
	var val interface{}
	var err error
	if metricType == "" || metricName == "" {
		return nil, runerr.ErrMetricNotFound
	}
	log.Println(metricType, metricName)
	switch metricType {
	case functionality.GaugeType:
		val, err = v.storage.GetGaugeByName(metricName)
		if err != nil {
			return nil, err
		}
	case functionality.CounterType:
		val, err = v.storage.GetCounterByName(metricName)
		if err != nil {
			return nil, err
		}
	default:
		return nil, runerr.ErrMetricNotFound
	}
	return val, nil

}

func NewValue(s storage.UpdateRepo) Value {
	return Value{s}
}
