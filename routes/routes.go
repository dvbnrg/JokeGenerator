package routes

import (
	"net/http"

	handle "github.com/dvbnrg/JokeGenerator/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handle.Index,
	},
	Route{
		"Generate",
		"GET",
		"/generate",
		handle.Generate,
	},
}
