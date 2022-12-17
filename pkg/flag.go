package pkg

import (
	"flag"
	"fmt"
)

func IsFlagPassed(args []string) (bool, error) {
	found := false
	i := 0
	flag.Visit(func(f *flag.Flag) {
		if f.Name == args[i] {
			found = true
		} else {
			i++
		}
	})

	return found, fmt.Errorf("missing %s arguments, use -help flag", args[i])
}
