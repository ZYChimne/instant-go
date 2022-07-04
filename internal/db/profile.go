package database

import (
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserInfo(user *model.User) error {
	return mongoDB.Users.FindOne(ctx, bson.M{"_id": user.UserID}).Decode(&user)
}
