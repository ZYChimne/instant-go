package api_test

import (
	"encoding/json"
	"zychimne/instant/config"
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/router"

	"github.com/gin-gonic/gin"
)

const APIVersion = "v1"

// Default Values
const DefaultEmail = "zychimne@instant.com"
const DefaultPhone = "1234567890"
const DefaultPassword = "Instant123@"
const DefaultUsername = "zychimne"

var r *gin.Engine

func init() {
	config.LoadConfig("../../config/dev.yml")
	database.ConnectPostgres()
	database.ConnectRedis()
	r = router.Create(false)
}

func FormatURL(prefix string, url string) string {
	return "/" + APIVersion + "/" + prefix + url
}

func MapToJson(m map[string]interface{}) string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

func ParseDataResponse(token []byte) interface{} {
	m := make(map[string]interface{})
	_ = json.Unmarshal(token, &m)
	return m["data"]
}
