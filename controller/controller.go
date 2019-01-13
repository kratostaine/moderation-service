package controller

import (
	"encoding/json"
	. "moderation-service/service"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Controller struct {
	Service
}

func (c *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal("Moderation service running fine")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) ValidateText(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentText := strings.ToLower(params["text"])

	err := c.Service.ValidateText(commentText)

	var responseMesage []byte
	if err == nil {
		w.WriteHeader(http.StatusOK)
		responseMesage, _ = json.Marshal("Comment text all good")
	} else {
		responseMesage, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(responseMesage)
	return
}
