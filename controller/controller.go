package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
}

func (c *Controller) getHealth(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal("All OK")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

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
		"GET",
		"/health",
		controller.getHealth,
	},
}

func NewRouter() *mux.Router {
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
