package helpers

import "time"

func GenerateDuration(millis int) time.Duration {
	return time.Duration(millis) * time.Millisecond
}
