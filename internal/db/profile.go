package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserProfile(userID string) *mongo.SingleResult {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil
	}
	return mongoDB.Users.FindOne(ctx, bson.M{"_id": oID})
}
