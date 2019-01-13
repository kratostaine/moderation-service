package controller

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
}

func (c *Controller) getHealth(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal("All OK")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
