package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var controller = &Controller{}

var routes = Routes{
	Route{
		"Health",
		http.MethodGet,
		"/health",
		controller.GetHealth,
	},
	Route{
		"Get data from repo",
		http.MethodGet,
		"/data",
		controller.GetData,
	},
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	log.Println("Registered routes:")
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name + ": " + route.Method + " " + route.Pattern)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
