package update

import (
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"sync"
)

type MemStorage struct {
	gauge      map[string]float64
	counter    map[string]int64
	gaugeMut   *sync.RWMutex
	counterMut *sync.RWMutex
}

func (m *MemStorage) GetAllMetrics() (map[string]float64, map[string]int64) {
	return m.gauge, m.counter
}

func (m *MemStorage) AddToGauge(name string, val float64) error {
	m.gaugeMut.Lock()
	m.gauge[name] = val
	m.gaugeMut.Unlock()
	return nil
}
func (m *MemStorage) GetGaugeByName(metricName string) (float64, error) {
	m.gaugeMut.RLock()
	val, ok := m.gauge[metricName]
	defer m.gaugeMut.RUnlock()
	if !ok {
		return 0.0, runerr.ErrMetricNotFound
	}
	return val, nil
}
func (m *MemStorage) GetCounterByName(metricName string) (int64, error) {
	m.counterMut.RLock()
	val, ok := m.counter[metricName]
	defer m.counterMut.RUnlock()
	if !ok {
		return 0, runerr.ErrMetricNotFound
	}
	return val, nil
}
func (m *MemStorage) AddToCounter(name string, val int64) error {
	_, ok := m.counter[name]
	m.counterMut.Lock()
	if !ok {
		m.counter[name] = val
	} else {
		m.counter[name] += val
	}
	m.counterMut.Unlock()

	return nil
}

func NewUpdateMemStorage() storage.UpdateRepo {
	return &MemStorage{
		make(map[string]float64),
		make(map[string]int64),
		&sync.RWMutex{},
		&sync.RWMutex{},
	}
}
