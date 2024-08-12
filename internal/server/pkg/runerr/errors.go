package runerr

import "errors"

var ErrEmptyMetricName = errors.New("empty metric name")
var ErrWrongMetricType = errors.New("wrong metric type")
var ErrWrongURL = errors.New("wrong url")
var ErrWrongHandlerType = errors.New("wrong handler type")
var ErrMetricNotFound = errors.New("metric not found")
