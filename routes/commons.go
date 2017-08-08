// Package routes defines all API's grouped by each file.
// TODO: Access Control
package routes

import "github.com/julienschmidt/httprouter"

// Route struct that stores info about a route
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler httprouter.Handle
	// TODO: Access Control
}
