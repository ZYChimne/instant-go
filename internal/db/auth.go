package database

import (
	"time"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(user model.User) (*mongo.InsertOneResult, error) {
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	return mongoDB.Users.InsertOne(
		ctx,
		bson.M{
			"mailbox":      user.MailBox,
			"phone":        user.Phone,
			"username":     user.Username,
			"password":     hash,
			"created":      time.Now(),
			"lastModified": time.Now(),
			"avatar":       user.Avatar,
			"gender":       user.Gender,
			"country":      user.Country,
			"province":     user.Province,
			"city":         user.City,
			"birthday":     user.Birthday,
			"school":       user.School,
			"company":      user.Company,
			"job":          user.Job,
			"myMode":       user.MyMode,
			"introduction": user.Introduction,
			"coverPhoto":   user.CoverPhoto,
			"tags":         user.Tags,
			"followings":   0,
			"followers":    0,
		},
	)
}

func GetUser(mailbox string, user *model.User) error {
	return mongoDB.Users.FindOne(ctx, bson.M{"mailbox": mailbox}).Decode(&user)
}
