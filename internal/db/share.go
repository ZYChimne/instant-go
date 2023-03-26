package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetShares(insID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(insID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Shares.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{
				{Key: "$match", Value: bson.M{"insID": oID}},
			},
			bson.D{{Key: "$sort", Value: bson.M{"_id": -1}}},
			bson.D{{Key: "$skip", Value: index}},
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "users",
					"localField":   "userID",
					"foreignField": "_id",
					"as":           "users",
					"pipeline": bson.A{
						bson.D{
							{
								Key:   "$project",
								Value: bson.M{"userID": 1, "username": 1, "avatar": 1},
							},
						},
					},
				},
			}},
			bson.D{
				{
					Key: "$replaceRoot",
					Value: bson.M{
						"newRoot": bson.M{
							"$mergeObjects": bson.A{bson.M{"$first": "$users"}, "$$ROOT"},
						},
					},
				},
			}},
		options.Aggregate().SetMaxTime(time.Second*2),
	)
}
