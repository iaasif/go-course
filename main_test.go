package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/events", createEvent)
	return r
}

func TestCreateEventSuccess(t *testing.T) {
	router := setupRouter()
	body := []byte(`{"name":"Go Meetup","description":"Learn Go","location":"Dhaka","dateTime":"2026-03-13T15:30:00Z"}`)

	req, _ := http.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusCreated, res.Code, res.Body.String())
	}
}

func TestCreateEventBadRequest(t *testing.T) {
	router := setupRouter()
	body := []byte(`{}`)

	req, _ := http.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusBadRequest, res.Code, res.Body.String())
	}
}
