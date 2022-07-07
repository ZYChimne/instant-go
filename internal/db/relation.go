package database

import (
	"errors"
	"strconv"
	"strings"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFollowing(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Followings.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func GetPotentialFollowing(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Followings.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func AddFollowing(following model.Following) error {
	session, err := mongoDB.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		userOID, err := primitive.ObjectIDFromHex(following.UserID)
		if err != nil {
			return nil, err
		}
		followingOID, err := primitive.ObjectIDFromHex(following.FollowingID)
		if err != nil {
			return nil, err
		}
		res1, err := mongoDB.Followings.InsertOne(ctx, bson.M{"userID": userOID, "followingID": followingOID, "created": following.Created})
		if err != nil {
			return res1, err
		}
		res2, err := mongoDB.Users.UpdateOne(ctx, bson.M{"_id": userOID}, bson.M{"$inc": bson.M{"followings": 1}})
		if err != nil {
			return res2, err
		}
		res3, err := mongoDB.Users.UpdateOne(ctx, bson.M{"_id": followingOID}, bson.M{"$inc": bson.M{"followers": 1}})
		if err != nil {
			return res3, err
		}
		if res2.ModifiedCount == 1 && res3.ModifiedCount == 1 {
			return nil, nil
		}
		return nil, errors.New(strings.Join([]string{"inc followings:", strconv.FormatInt(res2.ModifiedCount, 10), "inc followers:", strconv.FormatInt(res3.ModifiedCount, 10)}, " "))
	}
	_, err = session.WithTransaction(ctx, callback)
	return err
}

func RemoveFollowing(following model.Following) error {
	session, err := mongoDB.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		userOID, err := primitive.ObjectIDFromHex(following.UserID)
		if err != nil {
			return nil, err
		}
		followingOID, err := primitive.ObjectIDFromHex(following.FollowingID)
		if err != nil {
			return nil, err
		}
		res1, err := mongoDB.Followings.DeleteOne(ctx, bson.M{"userID": userOID, "followingID": followingOID})
		if err != nil {
			return res1, err
		}
		res2, err := mongoDB.Users.UpdateOne(ctx, bson.M{"_id": userOID}, bson.M{"$inc": bson.M{"followings": -1}})
		if err != nil {
			return res2, err
		}
		res3, err := mongoDB.Users.UpdateOne(ctx, bson.M{"_id": followingOID}, bson.M{"$inc": bson.M{"followers": -1}})
		if err != nil {
			return res3, err
		}
		if res1.DeletedCount == 1 && res2.ModifiedCount == 1 &&res3.ModifiedCount ==1 {
			return nil, nil
		}
		return nil, errors.New(strings.Join([]string{"delete:", strconv.FormatInt(res1.DeletedCount, 10), "inc followings:", strconv.FormatInt(res2.ModifiedCount, 10), "inc followers:", strconv.FormatInt(res3.ModifiedCount, 10)}, " "))
	}
	_, err = session.WithTransaction(ctx, callback)
	return err
}
