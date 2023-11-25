package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"zychimne/instant/tools"
	"zychimne/instant/pkg/schema"

	"github.com/stretchr/testify/assert"
)

var userID uint

func TestCreateAccount(t *testing.T) {
	// test creating an account
	w := httptest.NewRecorder()
	payload, _ := json.Marshal(schema.UpsertAccountRequest{
		Email:        DefaultEmail + "n",
		Phone:        DefaultPhone + "0",
		Password:     DefaultPassword,
		Username:     DefaultUsername + "a",
		Nickname:     DefaultUsername,
		Avatar:       "1",
		Gender:       0,
		Country:      "United States",
		State:        "Texas",
		City:         "Houston",
		ZipCode:      "77005",
		Birthday:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		School:       "Rice University",
		Company:      "Instant",
		Job:          "Software Engineer",
		MyMode:       "Chill",
		Introduction: "Keep your chin up!",
		CoverPhoto:   "1",
	})
	req, _ := http.NewRequest("POST", FormatURL("account", ""), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	userID = uint(ParseDataResponse(w.Body.Bytes()).(float64))
}

func TestGetAccount(t *testing.T) {
	// test getting an account
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("account", ""), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(userID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteAccount(t *testing.T) {
	// test deleting an account
	if userID == 0 {
		t.Errorf("userID is not set")
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", FormatURL("account", ""), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(userID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSearchAccount(t *testing.T) {
	// test searching an account
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("account", "/search?keyword=Evan&offset=0&limit=10"), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(userID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
