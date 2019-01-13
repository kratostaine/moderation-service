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

// swagger:route GET /health getHealth
//
// Gets the health of the service
//
//     Responses:
//       200:		healthy

func (c *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal("Moderation service running fine")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// swagger:route GET /validation/{text} validateText
//
// Validates if the text provided is objectionable or not
//
//     Responses:
//       200: 		text doesn't contain objectionable data
//       400: 		text contains objectionable data

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
