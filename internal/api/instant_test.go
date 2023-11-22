package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/schema"

	"github.com/stretchr/testify/assert"
)

func TestGetInstants(t *testing.T) {
	w := httptest.NewRecorder()
	payload, _ := json.Marshal(schema.UpsertInstantRequest{
		InstantType: 0,
		Content:     "Hello World!",
		Visibility:  0,
	})
	req, _ := http.NewRequest("POST", FormatURL("instant", ""), bytes.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(userID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteInstant(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", FormatURL("instant", "?instantID=6"), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(userID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
