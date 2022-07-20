package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserProfileByID(userID string) *mongo.SingleResult {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil
	}
	return mongoDB.Users.FindOne(ctx, bson.M{"_id": oID})
}

func QueryUsers(
	keyword string,
	index int64,
	pageSize int64,
) (*mongo.Cursor, error) {
	return mongoDB.Users.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{
				{
					Key: "$match",
					Value: bson.M{
						"$or": bson.A{
							bson.M{"mailbox": keyword},
							bson.M{"phone": keyword},
							bson.M{"$text": bson.M{"$search": keyword}},
						},
					},
				},
			},
			bson.D{{Key: "$sort", Value: bson.M{"score": bson.M{"$meta": "textScore"}}}},
			bson.D{{Key: "$skip", Value: index}},
			bson.D{{Key: "$limit", Value: pageSize}}},
		options.Aggregate().SetMaxTime(time.Second*2),
	)
}
