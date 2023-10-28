package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zychimne/instant/internal/api"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/schema"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	// test invalid email
	w := httptest.NewRecorder()
	payload, _ := json.Marshal(schema.LoginAccountRequest{
		Email:    DefaultEmail + "@",
		Password: DefaultPassword,
	})
	req, _ := http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, MapToJson(map[string]interface{}{
		"message": api.LoginError,
	}), w.Body.String())
	// test invalid phone
	w = httptest.NewRecorder()
	payload, _ = json.Marshal(schema.LoginAccountRequest{
		Phone:    DefaultPhone + "!",
		Password: DefaultPassword,
	})
	req, _ = http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, MapToJson(map[string]interface{}{
		"message": api.LoginError,
	}), w.Body.String())
	// test invalid password with email login
	w = httptest.NewRecorder()
	payload, _ = json.Marshal(schema.LoginAccountRequest{
		Email:    DefaultEmail,
		Password: "!" + DefaultPassword,
	})
	req, _ = http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, MapToJson(map[string]interface{}{
		"message": api.LoginError,
	}), w.Body.String())
	// test invalid password with phone login
	w = httptest.NewRecorder()
	payload, _ = json.Marshal(schema.LoginAccountRequest{
		Phone:    DefaultPhone,
		Password: "!" + DefaultPassword,
	})
	req, _ = http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, MapToJson(map[string]interface{}{
		"message": api.LoginError,
	}), w.Body.String())
	// test invalid login
	w = httptest.NewRecorder()
	payload, _ = json.Marshal(schema.LoginAccountRequest{
		Password: DefaultPassword,
	})
	req, _ = http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	// test email login
	w = httptest.NewRecorder()
	payload, _ = json.Marshal(schema.LoginAccountRequest{
		Email:    DefaultEmail,
		Password: DefaultPassword,
	})
	req, _ = http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	_, err := util.VerifyJwt(ParseDataResponse(w.Body.Bytes()).(string))
	assert.Nil(t, err)
	// test phone login
	w = httptest.NewRecorder()
	payload, _ = json.Marshal(schema.LoginAccountRequest{
		Phone:    DefaultPhone,
		Password: DefaultPassword,
	})
	req, _ = http.NewRequest("POST", FormatURL("auth", "/token"), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	_, err = util.VerifyJwt(ParseDataResponse(w.Body.Bytes()).(string))
	assert.Nil(t, err)
}
