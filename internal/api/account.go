package api

import (
	"errors"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var userSchema schema.UpsertAccountRequest
	if err := c.Bind(&userSchema); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(CreateAccountError))
		return
	}
	if !util.CheckStrongPassword(userSchema.Password) {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(CreateAccountError))
		return
	}
	if _, err := mail.ParseAddress(userSchema.Email); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LoginError))
		return
	}
	for _, char := range userSchema.Phone {
		if char < '0' || char > '9' {
			c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LoginError))
			return
		}
	}
	hash, err := util.HashPassword(userSchema.Password)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(CreateAccountError))
		return
	}
	if err := database.CheckGeo(userSchema.Country, userSchema.State, userSchema.City, userSchema.ZipCode); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(CreateAccountError))
		return
	}
	userModel := model.User{
		Email:          userSchema.Email,
		Phone:          userSchema.Phone,
		Password:       hash,
		Username:       userSchema.Username,
		Nickname:       userSchema.Nickname,
		UserType:       0,
		Avatar:         userSchema.Avatar,
		Gender:         userSchema.Gender,
		Country:        userSchema.Country,
		State:          userSchema.State,
		City:           userSchema.City,
		ZipCode:        userSchema.ZipCode,
		Birthday:       userSchema.Birthday,
		School:         userSchema.School,
		Company:        userSchema.Company,
		Job:            userSchema.Job,
		MyMode:         userSchema.MyMode,
		Introduction:   userSchema.Introduction,
		CoverPhoto:     userSchema.CoverPhoto,
		FollowingCount: 0,
		FollowerCount:  0,
	}
	err = database.CreateAccount(&userModel)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(CreateAccountError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": userModel.ID})
}

func DeleteAccount(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	var user model.User
	user.ID = userID
	err := database.DeleteAccountByID(&user)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(DeleteAccountError))
		return
	}
	c.Status(http.StatusOK)
}

func GetAccount(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	targetID := c.Query("userID")
	var user schema.AccountResponse
	if len(targetID) > 0 {
		targetUInt, err := strconv.ParseUint(targetID, 10, 64)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetAccountError))
			return
		}
		if err = database.GetAccountByID(uint(targetUInt), &user); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(GetAccountError))
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
		return
	}
	targetEmail := c.Query("email")
	if len(targetEmail) > 0 {
		if _, err := mail.ParseAddress(targetEmail); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetAccountError))
			return
		}
		if err := database.GetAccountByEmail(targetEmail, &user); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(GetAccountError))
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
		return
	}
	targetPhone := c.Query("phone")
	if len(targetPhone) > 0 {
		for _, char := range targetPhone {
			if char < '0' || char > '9' {
				c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetAccountError))
				return
			}
		}
		if err := database.GetAccountByPhone(targetPhone, &user); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(GetAccountError))
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
		return
	}
	targetUsername := c.Query("username")
	if len(targetUsername) > 0 {
		if err := database.GetAccountByUsername(targetUsername, &user); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(GetAccountError))
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
		return
	}
	database.GetAccountByID(userID, &user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func SearchAccounts(c *gin.Context) {
	_ = c.MustGet("UserID").(uint)
	keyword := c.Query("keyword")
	if len(keyword) == 0 {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(SearchAccountsError))
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 0, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(SearchAccountsError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 0, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(SearchAccountsError))
		return
	}
	var users []schema.BasicAccountResponse // TODO add relationship
	err = database.SearchAccounts(keyword, int(offset), int(limit), &users)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(SearchAccountsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CheckIfAccountExists(c *gin.Context) {
	email := c.Query("email")
	var exists int64
	if len(email) > 0 {
		if _, err := mail.ParseAddress(email); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusUnprocessableEntity, errors.New(CheckIfAccountExistsError))
			return
		}
		err := database.CheckIfAccountExistsByEmail(email, &exists)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(CheckIfAccountExistsError))
			return
		}
		if exists > 0 {
			c.JSON(http.StatusOK, gin.H{"data": 1})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": 0})
		return
	}
	phone := c.Query("phone")
	if len(phone) > 0 {
		for _, char := range phone {
			if char < '0' || char > '9' {
				c.AbortWithError(http.StatusUnprocessableEntity, errors.New(CheckIfAccountExistsError))
				return
			}
		}
		err := database.CheckIfAccountExistsByPhone(phone, &exists)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(CheckIfAccountExistsError))
			return
		}
		if exists > 0 {
			c.JSON(http.StatusOK, gin.H{"data": 1})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": 0})
		return
	}
	username := c.Query("username")
	if len(username) > 0 {
		err := database.CheckIfAccountExistsByUsername(username, &exists)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(CheckIfAccountExistsError))
			return
		}
		if exists > 0 {
			c.JSON(http.StatusOK, gin.H{"data": 1})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": 0})
		return
	}
}
