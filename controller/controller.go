package controller

import (
	"encoding/json"
	. "moderation-service/service"
	"net/http"
)

type Controller struct {
	Service
}

func (c *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal("All OK")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) GetData(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(c.Service.GetData())
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
