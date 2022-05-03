package handlr

// Aliases the RouteFunc method from Handlr.Router to Handlr.
func (h *Handlr) RouteFunc(path string, routeHandler RouteHandlerFunc) {
	h.router.RouteFunc(path, routeHandler)
}

// Aliases the HandleFunc method from Handlr.Router to Handlr.
func (h *Handlr) HandleFunc(path string, actionHandler ActionHandlerFunc) {
	h.router.HandleFunc(path, actionHandler)
}
