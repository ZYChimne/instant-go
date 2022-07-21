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

func QueryUsers(userID string, keyword string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
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
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{Key: "$project", Value: bson.M{"username": 1, "avatar": 1}}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "followings",
					"localField":   "_id",
					"foreignField": "followingID",
					"as":           "followings",
					"pipeline": bson.A{
						bson.D{
							{Key: "$match", Value: bson.M{"userID": oID}},
						},
						bson.D{
							{Key: "$project", Value: bson.M{"isFriend": 1}},
						},
					},
				},
			}},
			bson.D{
				{
					Key: "$project",
					Value: bson.M{
						"username": 1,
						"avatar":   1,
						"isFollowing": bson.M{
							"$cond": bson.M{
								"if":   bson.M{"$gt": bson.A{bson.M{"$size": "$followings"}, 0}},
								"then": true,
								"else": false,
							},
						},
						"isFriend": bson.M{"$first": "$followings.isFriend"},
					},
				},
			},
		},
		options.Aggregate().SetMaxTime(time.Second*2),
	)
}
