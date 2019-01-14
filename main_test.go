package main

import (
	"moderation-service/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestGetHealth(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expectedBody := "\"Moderation service running fine\""

	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected - Moderation service running fine. Got %s", body)
	}
}

func TestValidateTextWithObjectionableContent(t *testing.T) {
	req, _ := http.NewRequest("GET", "/validation/kill%20the%20person", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	expectedBody := "\"Objectionable data found\""

	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected - %s. Got %s", expectedBody, body)
	}
}

func TestValidateTextWithNoObjectionableContent(t *testing.T) {
	req, _ := http.NewRequest("GET", "/validation/welcome%20the%20person", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expectedBody := "\"Comment text all good\""

	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected - %s. Got %s", expectedBody, body)
	}
}

func TestValidateTextWithNoObjectionableContentAndFalsePositive(t *testing.T) {
	req, _ := http.NewRequest("GET", "/validation/welcome%20dikshit%20to%20home", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expectedBody := "\"Comment text all good\""

	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected - %s. Got %s", expectedBody, body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router = controller.Router()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
