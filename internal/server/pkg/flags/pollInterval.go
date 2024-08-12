package flags

import (
	"fmt"
	"strconv"
)

type PollInterval struct {
	Interval int
}

func (a *PollInterval) String() string {
	return fmt.Sprintf("%d", a.Interval)
}

func (a *PollInterval) Type() string {
	return "int"
}
func (a *PollInterval) Set(flagValue string) error {
	parseInt, err := strconv.ParseInt(flagValue, 10, 64)
	if err != nil {
		return err
	}
	a.Interval = int(parseInt)
	return nil
}

func GetPollInterval() (*PollInterval, error) {
	a := &PollInterval{
		Interval: 2,
	}
	err := setEnvFlag(a, "POLL_INTERVAL", "p", "p", "consuming data frequency in seconds")
	if err != nil {
		return nil, err
	}

	return a, nil
}
