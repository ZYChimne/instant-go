package database

import (
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFollowing(userID string, index int64, pageSize int64) (*mongo.Cursor, error)  {
	return mongoDB.Following.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func GetPotentialFollowing(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Following.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func AddFollowing(follower model.Follower) (*mongo.InsertOneResult, error) {
	return mongoDB.Following.InsertOne(ctx, bson.M{"userID":follower.UserID, "followingID":follower.FollowingID, "created":follower.Created})
}

func RemoveFollowing(follower model.Follower) (*mongo.DeleteResult, error) {
	return mongoDB.Following.DeleteOne(ctx, bson.M{"userID": follower.UserID, "followingID": follower.FollowingID})
}
