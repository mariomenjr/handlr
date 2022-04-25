package handlr

import (
	"strings"
)

// Trims slashes from a string.
func trimSlash(endpoint string) string {
	return strings.Trim(endpoint, "/")
}
