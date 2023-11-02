package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountries(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("geo", "/countries"), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetStates(t *testing.T) {
	// without country id
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("geo", "/states"), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	// with country id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", FormatURL("geo", "/states?cID=1"), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCities(t *testing.T) {
	// without state id
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("geo", "/cities"), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	// with state id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", FormatURL("geo", "/cities?sID=1"), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
