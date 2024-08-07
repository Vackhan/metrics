package storage

type UpdateRepo interface {
	AddToGauge(name string, val float64) error
	AddToCounter(name string, val int64) error
}
