package storage

type UpdateRepo interface {
	AddToGauge(name string, val float64) error
	AddToCounter(name string, val int64) error
	GetGaugeByName(metricName string) (float64, error)
	GetCounterByName(metricName string) (int64, error)
	GetAllMetrics() (map[string]float64, map[string]int64)
}
