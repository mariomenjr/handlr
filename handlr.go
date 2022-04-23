package handlr

import (
	"fmt"
	"net/http"
)

// Allows end user to create an instance of Handlr
func New() *Handlr {
	return &Handlr{http.NewServeMux(), Router{path: "/"}}
}

// It houses the main Router as well as the mux instances.
type Handlr struct {
	mux    *http.ServeMux
	router Router
}

// Registers Routers and ListenAndServer over Handlr.mux
func (h *Handlr) Start(portNumber int) error {
	h.router.regiterRoutesAndHandler(h.mux)
	return h.listenAndServe(portNumber)
}

// ListenAndServer over Handlr.mux
func (h *Handlr) listenAndServe(portNumber int) error {
	portString := fmt.Sprintf(":%d", portNumber)

	fmt.Printf("> Server started on port %s", portString)

	return http.ListenAndServe(portString, h.mux)
}
