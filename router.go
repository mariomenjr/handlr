package handlr

import (
	"net/http"
	"path"
	"strings"
)

// A router instance will house paths and handlers.
// It will also keep track of their hierarchy.
//
// Ideally, you'd like to distribute your handlers
// into separate files and plugin them in as Routes.
type Router struct {
	path     string
	parent   *Router
	children []*Router
	handler  *ActionHandler
}

// Allows Route registration.
// You don't program behavior through this method.
func (rt *Router) Route(path string, routeHandler RouteHandler) {
	router := &Router{path: path, parent: rt}
	routeHandler(router)

	rt.children = append(rt.children, router)
}

// Allows Handler registration which gives you the ability
// to tie a behavior to a path.
// i.e. Get a record from database by hiting URL:
// 			http://example.org/get/record?id=1
func (rt *Router) Handler(path string, actionHandler ActionHandler) {
	router := &Router{path: path, parent: rt, handler: &actionHandler}

	rt.children = append(rt.children, router)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := rt.findHandler(r); handler != nil {
		(*handler)(w, r)
		return
	}
	http.NotFound(w, r)
}

// It attempts to produce Router's endpoint path
// based on parent and children routes.
func (rt *Router) buildPath() string {
	if rt.parent == nil {
		return rt.path
	}
	return path.Join(rt.parent.buildPath(), trimSlash(rt.path))
}

// Recursively find a handler to the request
func (rt *Router) findHandler(r *http.Request) *ActionHandler {
	for _, v := range rt.children {
		if v.handler == nil {
			continue
		}

		if v.isMatch(r) {
			return v.handler
		}

		v.findHandler(r)
	}
	return nil
}

// Asserts if Router path matches URL from Request
func (rt *Router) isMatch(r *http.Request) bool {
	rp := strings.Split(trimSlash(rt.buildPath()), "/")
	ep := strings.Split(trimSlash(r.URL.Path), "/")

	if len(rp) == len(ep) {
		m := true
		for i := 0; i < len(rp); i++ {
			rs := strings.ToLower(rp[i])
			es := strings.ToLower(ep[i])

			// Very basic slug matching. Regex later.
			m = m && rs == es || (rs[0] == ':' && es != "")
		}
		return m
	}
	return false
}
