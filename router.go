package handlr

import (
	"log"
	"net/http"
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
	handler  *func(w http.ResponseWriter, r *http.Request)
}

// Allows Route registration.
// You don't program behavior through this method.
func (r *Router) Route(path string, routeHandler func(r *Router)) {
	router := &Router{path: path, parent: r}
	routeHandler(router)

	r.children = append(r.children, router)
}

// Allows Handler registration which gives you the ability
// to tie a behavior to a path.
// i.e. Get a record from database by hiting URL:
// 			http://example.org/get/record?id=1
func (r *Router) Handler(path string, actionHandler func(w http.ResponseWriter, r *http.Request)) {
	router := &Router{path: path, parent: r, handler: &actionHandler}

	r.children = append(r.children, router)
}

// Recursively register handlers for paths.
// An error will be thrown if the same path is registered twice, no ServeMux
// instance is provided, or mux.HandleFunc throws an error itself.
func (r *Router) regiterRoutesAndHandler(mux *http.ServeMux) {
	if mux == nil {
		log.Fatal("router: No *http.ServeMux instance provided for registering routes and handlers.")
	}

	for _, v := range r.children {
		if v.handler != nil {
			mux.HandleFunc(v.buildPath(), *v.handler)
		}

		v.regiterRoutesAndHandler(mux)
	}
}
