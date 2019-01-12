package main

import (
	"log"
	"moderation-service/controller"
	"net/http"
)

func main() {
	router := controller.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
