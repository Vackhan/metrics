package update

import (
	"github.com/Vackhan/metrics/internal/server/pkg/functionality"
	"github.com/Vackhan/metrics/internal/server/pkg/runerr"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
	"log"
	"strconv"
	"strings"
)

type Update struct {
	storage storage.UpdateRepo
}

func (u *Update) DoUpdate(path string) error {
	urlData := strings.Split(path[1:], "/")
	if len(urlData) != 4 {
		return runerr.ErrWrongURL
	}
	log.Println(urlData)
	var err error
	switch urlData[1] {
	case functionality.GaugeType:
		var g float64
		if g, err = strconv.ParseFloat(urlData[3], 64); err != nil {
			return runerr.ErrWrongMetricType
		}

		err = u.storage.AddToGauge(urlData[2], g)
		if err != nil {
			return err
		}
	case functionality.CounterType:
		var c int64
		if c, err = strconv.ParseInt(urlData[3], 10, 64); err != nil {
			log.Println(err)
			return runerr.ErrWrongMetricType
		}
		err := u.storage.AddToCounter(urlData[2], c)
		if err != nil {
			return err
		}
	default:
		return runerr.ErrWrongMetricType
	}
	return nil
}

func NewUpdate(s storage.UpdateRepo) Update {
	return Update{s}
}
