package main

import (
	"log"
	"moderation-service/controller"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	a := App{}
	a.Router = controller.Router()
	a.Run(":9000")
}
