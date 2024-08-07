package update

import "github.com/Vackhan/metrics/internal/server/internal/storage"

type MemStorage struct {
	gauge   map[string]float64
	counter map[string][]int64
}

func (m *MemStorage) AddToGauge(name string, val float64) error {
	m.gauge[name] = val
	return nil
}
func (m *MemStorage) ImplementRepository() {}
func (m *MemStorage) AddToCounter(name string, val int64) error {
	v, ok := m.counter[name]
	if !ok {
		v = []int64{}
	}
	v = append(v, val)
	m.counter[name] = v
	return nil
}

func NewMemStorage() storage.UpdateRepo {
	return &MemStorage{make(map[string]float64), make(map[string][]int64)}
}
