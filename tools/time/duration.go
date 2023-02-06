package time

import "time"

func ParseDuration(str string) time.Duration {
	if dra, err := time.ParseDuration(str); err != nil {
		return time.Millisecond * 100
	} else {
		return dra
	}
}
