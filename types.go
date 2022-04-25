package handlr

import "net/http"

type ActionHandler func(w http.ResponseWriter, r *http.Request)
type RouteHandler func(r *Router)
