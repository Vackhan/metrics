package flags

import (
	"errors"
	"fmt"
	"strings"
)

type Address struct {
	Host string
	Port string
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func (a *Address) Type() string {
	return "string"
}
func (a *Address) Set(flagValue string) error {
	values := strings.Split(flagValue, ":")
	if len(values) != 2 {
		return errors.New("address should contain host and port, i.e. \"127.0.0.0:1010\", \":8080\" etc")
	}
	a.Host = values[0]
	a.Port = values[1]
	return nil
}

func GetAddress() (*Address, error) {
	a := &Address{
		Host: "",
		Port: "8080",
	}
	err := setEnvFlag(a, "ADDRESS", "a", "a", "host and port of the listener")
	if err != nil {
		return nil, err
	}

	return a, nil
}
