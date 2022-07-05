package database

import (
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFollowing(userID string, index int64, pageSize int64) (*mongo.Cursor, error)  {
	return mongoDB.Followings.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func GetPotentialFollowing(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Followings.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func AddFollowing(follower model.Following) (*mongo.InsertOneResult, error) {
	return mongoDB.Followings.InsertOne(ctx, bson.M{"userID":follower.UserID, "followingID":follower.FollowingID, "created":follower.Created})
}

func RemoveFollowing(follower model.Following) (*mongo.DeleteResult, error) {
	return mongoDB.Followings.DeleteOne(ctx, bson.M{"userID": follower.UserID, "followingID": follower.FollowingID})
}
