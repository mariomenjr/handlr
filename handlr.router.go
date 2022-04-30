package handlr

// Aliases the RouteFunc method from Handlr.Router to Handlr.
func (h *Handlr) RouteFunc(path string, routeHandler RouteHandler) {
	h.router.RouteFunc(path, routeHandler)
}

// Aliases the HandlerFunc method from Handlr.Router to Handlr.
func (h *Handlr) HandlerFunc(path string, actionHandler ActionHandler) {
	h.router.HandlerFunc(path, actionHandler)
}
