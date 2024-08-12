package flags

import (
	"fmt"
	"strconv"
)

type ReportInterval struct {
	Interval int
}

func (a *ReportInterval) String() string {
	return fmt.Sprintf("%d", a.Interval)
}

func (a *ReportInterval) Type() string {
	return "int"
}
func (a *ReportInterval) Set(flagValue string) error {
	parseInt, err := strconv.ParseInt(flagValue, 10, 64)
	if err != nil {
		return err
	}
	a.Interval = int(parseInt)
	return nil
}

func GetReportInterval() (*ReportInterval, error) {
	a := &ReportInterval{
		Interval: 10,
	}
	err := setEnvFlag(a, "REPORT_INTERVAL", "r", "r", "server send frequency in seconds")
	if err != nil {
		return nil, err
	}

	return a, nil
}
