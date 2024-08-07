package upd

import (
	"errors"
	"github.com/Vackhan/metrics/internal/server"
	"github.com/Vackhan/metrics/internal/server/internal/runerr"
	"github.com/Vackhan/metrics/internal/server/internal/storage"
	"strconv"
	"strings"
)

const gaugeType = "gauge"
const counterType = "counter"

var listType = []string{gaugeType, counterType}

type Update struct {
	storage server.Repository
}

func (u *Update) DoUpdate(path string) error {
	urlData := strings.Split(path[1:], "/")
	if len(urlData) != 4 {
		return runerr.ErrWrongURL
	}
	var err error
	updateRepo, ok := u.storage.(storage.UpdateRepo)
	if !ok {
		return errors.New("wrong repo")
	}
	switch urlData[1] {
	case gaugeType:
		var g float64
		if g, err = strconv.ParseFloat(urlData[3], 64); err != nil {
			return runerr.ErrWrongMetricType
		}

		err := updateRepo.AddToGauge(urlData[2], g)
		if err != nil {
			return err
		}
	case counterType:
		var c int64
		if c, err = strconv.ParseInt(urlData[3], 10, 64); err == nil {
			return runerr.ErrWrongMetricType
		}
		updateRepo.AddToCounter(urlData[2], c)
	default:
		return runerr.ErrWrongMetricType
	}
	return nil
}

func NewUpdate(s server.Repository) Update {
	return Update{s}
}
