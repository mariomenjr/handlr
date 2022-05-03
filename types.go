package handlr

import "net/http"

type ActionHandlerFunc func(w http.ResponseWriter, r *http.Request)
type RouteHandlerFunc func(r *Router)

type RouteHandler interface {
	RouteHTTP(r *Router)
}
