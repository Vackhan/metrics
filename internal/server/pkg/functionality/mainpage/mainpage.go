package mainpage

import (
	"github.com/Vackhan/metrics/internal/server/pkg/functionality"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
)

type MainPage struct {
	storage storage.UpdateRepo
}

func NewMainPage(storage storage.UpdateRepo) MainPage {
	return MainPage{storage: storage}
}

func (m MainPage) GetListOfMetrics() []Metric {
	gauge, counter := m.storage.GetAllMetrics()
	metrics := make([]Metric, 0, len(gauge)+len(counter))
	for name, val := range gauge {
		metrics = append(metrics, Metric{
			Type: functionality.GaugeType,
			Name: name,
			Val:  val,
		})
	}
	for name, val := range counter {
		metrics = append(metrics, Metric{
			Type: functionality.CounterType,
			Name: name,
			Val:  val,
		})
	}
	return metrics
}

type Metric struct {
	Type string
	Name string
	Val  any
}
