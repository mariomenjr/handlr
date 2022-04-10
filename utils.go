package handlr

import (
	"path"
	"strings"
)

// Given a Router instance and attempts to
// produce its endpoint path based on parent
// and children routes.
func (r *Router) buildPath() string {
	if r.parent == nil {
		return r.path
	}
	return path.Join(r.parent.buildPath(), trimSlash(r.path))
}

// Trims slashes from a string.
func trimSlash(endpoint string) string {
	return strings.Trim(endpoint, "/")
}
