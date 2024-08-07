package functionalityErrors

import "errors"

var ErrEmptyMetricName = errors.New("empty metric name")
var ErrWrongMetricType = errors.New("wrong metric type")
var ErrWrongUrl = errors.New("wrong url")
