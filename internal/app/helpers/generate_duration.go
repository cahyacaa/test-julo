package helpers

import "time"

func GenerateDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}
