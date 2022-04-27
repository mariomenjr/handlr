package handlr

// Aliases the Route method from Handlr.Router to Handlr.
func (h *Handlr) Route(path string, routeHandler RouteHandler) {
	h.router.Route(path, routeHandler)
}

// Aliases the Handler method from Handlr.Router to Handlr.
func (h *Handlr) Handler(path string, actionHandler ActionHandler) {
	h.router.Handler(path, actionHandler)
}
