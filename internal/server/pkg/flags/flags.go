package flags

import (
	"github.com/spf13/pflag"
	"os"
)

func setEnvFlag(a pflag.Value, envVal, flagName, flagShort, flagUsage string) error {
	addressEnv, ok := os.LookupEnv(envVal)
	if !ok {
		pflag.VarP(a, flagName, flagShort, flagUsage)
	} else {
		err := a.Set(addressEnv)
		if err != nil {
			return err
		}
	}
	return nil
}
