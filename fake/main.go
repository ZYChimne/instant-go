package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
	"zychimne/instant/config"
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/brianvoe/gofakeit/v6"
)

// General Configurations
const TestHttp = true
const GenerateFakeData = false
const ServerAddress = "http://localhost:8081"
const APIVersion = "v1"

// Specific Configurations
const MaxArraySize = 1000
const UserCount = 2000000 // 2 million, contains duplicates
// Runtime Configurations
var Token string
var Client *http.Client

// Default Values
const DefaultEmail = "zychimne@instant.com"
const DefaultPhone = "1234567890"
const DefaultPassword = "Instant123@"
const DefaultUsername = "zychimne"

func request(method string, url string, data []byte) {
	req, err := http.NewRequest(method, strings.Join([]string{ServerAddress, APIVersion, url}, "/"), bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	if len(Token) > 0 {
		req.Header.Set("Authorization", "Bearer "+Token)
	}
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(bytes))
}

func testRegister() {
	data, err := json.Marshal(schema.CreateUser{
		Email:        DefaultEmail,
		Phone:        DefaultPhone,
		Password:     DefaultPassword,
		Username:     DefaultUsername,
		Nickname:     "Evan Cheng",
		Avatar:       "0",
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
		Introduction: "SWE @ Instant",
		CoverPhoto:   "0",
	})
	if err != nil {
		panic(err)
	}
	request(http.MethodPost, "auth/register", data)
}

func testGetToken() {
	data, err := json.Marshal(schema.LoginUser{
		Email:    DefaultEmail,
		Password: DefaultPassword,
	})
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(strings.Join([]string{ServerAddress, APIVersion, "auth/token"}, "/"), "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	type TokenResponse struct {
		Token string `json:"data"`
	}
	var tokenResponse TokenResponse
	err = json.Unmarshal(bytes, &tokenResponse)
	if err != nil {
		panic(err)
	}
	println("Token", tokenResponse.Token)
	Token = tokenResponse.Token
}

func testAddFollowing() {
	data, err := json.Marshal(schema.UpdateFollowingRequest{
		TargetID: 2,
	})
	if err != nil {
		panic(err)
	}
	request(http.MethodPost, "relation", data)
}

func testRemoveFollowing() {
	data, err := json.Marshal(schema.UpdateFollowingRequest{
		TargetID: 2,
	})
	if err != nil {
		panic(err)
	}
	request(http.MethodDelete, "relation", data)
}

func fakeUsers() {
	if UserCount%MaxArraySize != 0 {
		panic("UserCount must be a multiple of " + string(rune(MaxArraySize)))
	}
	users := make([]model.User, MaxArraySize)
	hash, err := util.HashPassword(DefaultPassword)
	if err != nil {
		panic(err)
	}
	total := UserCount / MaxArraySize
	start := time.Now()
	last := start
	for i := 0; i < total; i++ {
		for j := 0; j < MaxArraySize; j++ {
			addr := gofakeit.Address()
			users[j] = model.User{
				Email:          gofakeit.Email(),
				Phone:          gofakeit.Phone(),
				Password:       hash,
				Username:       gofakeit.Username()+string(i),
				Nickname:       gofakeit.Name(),
				Type:           1,
				Avatar:         "1",
				Gender:         0,
				Country:        addr.Country,
				State:          addr.State,
				City:           addr.City,
				ZipCode:        addr.Zip,
				Birthday:       gofakeit.Date(),
				School:         "Rice University",
				Company:        gofakeit.Company(),
				Job:            gofakeit.JobTitle(),
				MyMode:         gofakeit.AdjectiveDescriptive(),
				Introduction:   gofakeit.Sentence(10),
				CoverPhoto:     "1",
				FollowingCount: 0,
				FollowerCount:  0,
			}
		}
		err = database.CreateUsers(&users)
		if err != nil {
			println(err)
		}
		println("Faked ", i*MaxArraySize, " users", time.Since(last).String())
		last = time.Now()
	}
	println("Faking ", UserCount, " users: ", time.Since(start).String())
}

func main() {
	config.LoadConfig()
	database.ConnectPostgres()
	Client = &http.Client{}
	if TestHttp {
		// testRegister()
		testGetToken()
		testAddFollowing()
		testRemoveFollowing()
	}
	if GenerateFakeData {
		fakeUsers()
	}
}
