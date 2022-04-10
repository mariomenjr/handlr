package handlr

import "net/http"

// Aliases the Route method from Handlr.Router to Handlr.
func (h *Handlr) Route(path string, routeHandler func(r *Router)) {
	h.router.Route(path, routeHandler)
}

// Aliases the Handler method from Handlr.Router to Handlr.
func (h *Handlr) Handler(path string, actionHandler func(w http.ResponseWriter, r *http.Request)) {
	h.router.Handler(path, actionHandler)
}
