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

const followerID = 1
const followingID = 2

func TestFollow(t *testing.T) {
	// test follow
	w := httptest.NewRecorder()
	payload, _ := json.Marshal(schema.UpsertFollowingRequest{
		TargetID: followingID,
	})
	req, _ := http.NewRequest("POST", FormatURL("relation", ""), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(followerID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetFollowings(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("relation", "/following?offset=0&limit=10"), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(followerID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFollowers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("relation", "/follower?offset=0&limit=10"), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(followingID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFriends(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("relation", "/friend?offset=0&limit=10"), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(followerID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPotentialFollowing(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", FormatURL("relation", "/potential?offset=0&limit=10"), nil)
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(followerID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUnfollow(t *testing.T) {
	// test unfollow
	w := httptest.NewRecorder()
	payload, _ := json.Marshal(schema.UpsertFollowingRequest{
		TargetID: followingID,
	})
	req, _ := http.NewRequest("DELETE", FormatURL("relation", ""), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+util.GenerateJwt(followerID))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
