package upd

import (
	"github.com/Vackhan/metrics/internal/server/internal/runerr"
	"github.com/Vackhan/metrics/internal/server/internal/storage"
	"log"
	"strconv"
	"strings"
)

const gaugeType = "gauge"
const counterType = "counter"

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
	case gaugeType:
		var g float64
		if g, err = strconv.ParseFloat(urlData[3], 64); err != nil {
			return runerr.ErrWrongMetricType
		}

		err = u.storage.AddToGauge(urlData[2], g)
		if err != nil {
			return err
		}
	case counterType:
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
