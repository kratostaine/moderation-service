package main

import (
	"log"
	"moderation-service/controller"
	"net/http"
)

func main() {
	router := controller.Router()
	log.Fatal(http.ListenAndServe(":8000", router))
}
