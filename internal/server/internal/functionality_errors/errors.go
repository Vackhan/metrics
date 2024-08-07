package functionality_errors

import "errors"

var EmptyMetricName = errors.New("empty metric name")
var WrongMetricType = errors.New("wrong metric type")
var WrongUrl = errors.New("wrong url")
