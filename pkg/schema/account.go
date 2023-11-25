package schema

import (
	"time"
)

type UpsertAccountRequest struct {
	Email        string    `json:"email" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
	Password     string    `json:"password"`
	Username     string    `json:"username"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Gender       int       `json:"gender"`
	Country      string    `json:"country"`
	State        string    `json:"state"`
	City         string    `json:"city"`
	ZipCode      string    `json:"zipCode"`
	Birthday     time.Time `json:"birthday"`
	School       string    `json:"school"`
	Company      string    `json:"company"`
	Job          string    `json:"job"`
	MyMode       string    `json:"myMode"`
	Introduction string    `json:"intro"`
	CoverPhoto   string    `json:"coverPhoto"`
}

type LoginAccountRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}

type AccountResponse struct {
	ID           uint      `json:"userID"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Username     string    `json:"username" `
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Gender       int       `json:"gender"`
	Country      string    `json:"country"`
	State        string    `json:"state"`
	City         string    `json:"city"`
	ZipCode      string    `json:"zipCode"`
	Birthday     time.Time `json:"birthday" binding:"required"`
	School       string    `json:"school"`
	Company      string    `json:"company"`
	Job          string    `json:"job"`
	MyMode       string    `json:"myMode"`
	Introduction string    `json:"intro"`
	CoverPhoto   string    `json:"coverPhoto"`
	Followings   int       `json:"followings"`
	Followers    int       `json:"followers"`
}

type BasicAccountResponse struct {
	ID       uint   `json:"userID"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	UserType int    `json:"userType"`
}

type QueryAccountResponse struct {
	BasicAccountResponse
	IsFollowing bool `json:"isFollowing"`
	IsFriend    bool `json:"isFriend"`
}
