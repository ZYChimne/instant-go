package database

import (
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserInfo(user *model.User) error {
	oID, err := primitive.ObjectIDFromHex(user.UserID)
	if err != nil {
		return nil
	}
	return mongoDB.Users.FindOne(ctx, bson.M{"_id": oID}).Decode(&user)
}
