package helpers

import "strings"

func GenerateKey(key ...string) string {
	return strings.Join(key, ":")
}
