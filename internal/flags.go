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

		// check if it’s a flag (starts with -)
		if len(arg) > 0 && arg[0] == '-' {
			// remove leading '-' or '--'
			key := arg
			for len(key) > 0 && key[0] == '-' {
				key = key[1:]
			}

			// assign value if present
			if i+1 < len(os.Args) && (len(os.Args[i+1]) == 0 || os.Args[i+1][0] != '-') {
				flags[key] = os.Args[i+1]
				i++ // skip value
			} else {
				flags[key] = ""
			}
		}
	}
	return flags
}
