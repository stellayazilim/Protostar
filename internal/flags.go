package internal

import (
	"os"
	"time"
)

type Flag map[string]string

func (f Flag) AddDefaultFlags() Flag {

	// defaults
	if _, ok := f["date"]; !ok {
		f["date"] = time.Now().Format("2006-01-02")
	}
	if _, ok := f["output"]; !ok {
		f["output"] = "CHANGELOG.md"
	}

	return f
}

func ParseFlags() Flag {
	flags := make(Flag)
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if len(arg) > 1 && arg[0] == '-' {
			key := arg[2:]
			if i+1 < len(os.Args) {
				flags[key] = os.Args[i+1]
				i++ // skip next arg as it’s the value
			} else {
				flags[key] = ""
			}
		}
	}

	return flags
}
