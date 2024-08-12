package flags

import (
	"fmt"
	"strconv"
	"time"
)

type TTL struct {
	Duration time.Duration
}

func (a *TTL) String() string {
	return fmt.Sprintf("%d", a.Duration)
}

func (a *TTL) Type() string {
	return "duration"
}
func (a *TTL) Set(flagValue string) error {
	parseInt, err := strconv.ParseInt(flagValue, 10, 64)
	if err != nil {
		return err
	}
	a.Duration = time.Duration(parseInt)
	return nil
}

func GetTTL() (*TTL, error) {
	a := &TTL{
		Duration: 0,
	}
	err := setEnvFlag(a, "TTL", "ttl", "t", "time to live for agent")
	if err != nil {
		return nil, err
	}

	return a, nil
}
