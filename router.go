package main

import (
	"github.com/escherpad/api-server/routes"
	"github.com/julienschmidt/httprouter"
)

// NewRouter Returns a new router with all handlers attached
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	deployAPI(router, [][]routes.Route{routes.AuthRoutes})

	return router
}

func deployAPI(router *httprouter.Router, routes [][]routes.Route) {
	for _, rs := range routes {
		for _, r := range rs {
			router.Handle(r.Method, r.Pattern, r.Handler)
		}
	}
}