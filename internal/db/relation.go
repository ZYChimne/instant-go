package database

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFollowings(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Followings.Find(
		ctx,
		bson.M{"userID": oID},
		options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize),
	)
}

func GetFollowers(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Followings.Find(
		ctx,
		bson.M{"followingID": oID},
		options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize),
	)
}

// Find people who follow me yet I don't follow them
func GetPotentialFollowings(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Followings.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.M{"followingID": oID, "isFriend": false}}},
			bson.D{{Key: "$sort", Value: bson.M{"lastModified": -1}}},
			bson.D{{Key: "$skip", Value: index}},
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "users",
					"localField":   "difference",
					"foreignField": "_id",
					"as":           "user",
				},
			}},
			bson.D{{Key: "$replaceRoot", Value: bson.M{"newRoot": bson.M{"$first": "$user"}}}},
		},
		options.Aggregate().SetMaxTime(2*time.Second),
	)
}

// ONLY FOR DEBUG AND TEST
func GetAllUsers(index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Users.Find(
		ctx,
		bson.M{},
		options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize),
	)
}

func AddFollowing(following model.Following) error {
	if following.UserID == following.FollowingID {
		return errors.New("userID and followingID must not be the same")
	}
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
		now := time.Now()
		res1, err := mongoDB.Followings.UpdateOne(
			ctx,
			bson.M{"followingID": userOID, "userID": followingOID},
			bson.M{"$set": bson.M{"lastModified": now, "isFriend": true}},
		)
		if err != nil {
			return res1, err
		}
		res2, err := mongoDB.Followings.InsertOne(
			ctx,
			bson.M{
				"userID":       userOID,
				"followingID":  followingOID,
				"isFriend":     res1.ModifiedCount==1,
				"created":      now,
				"lastModified": now,
			},
		)
		if err != nil {
			return res2, err
		}
		res3, err := mongoDB.Users.UpdateOne(
			ctx,
			bson.M{"_id": userOID},
			bson.M{"$inc": bson.M{"followings": 1}},
		)
		if err != nil {
			return res3, err
		}
		res4, err := mongoDB.Users.UpdateOne(
			ctx,
			bson.M{"_id": followingOID},
			bson.M{"$inc": bson.M{"followers": 1}},
		)
		if err != nil {
			return res4, err
		}
		if res3.ModifiedCount != 1 || res4.ModifiedCount != 1 {
			return nil, errors.New(
				strings.Join(
					[]string{
						"inc followings:",
						strconv.FormatInt(res3.ModifiedCount, 10),
						"inc followers:",
						strconv.FormatInt(res4.ModifiedCount, 10),
					},
					" ",
				),
			)
		}
		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	return err
}

func RemoveFollowing(following model.Following) error {
	if following.UserID == following.FollowingID {
		return errors.New("userID and followingID must not be the same")
	}
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
		res1, err := mongoDB.Followings.UpdateOne(
			ctx,
			bson.M{"followingID": userOID, "userID": followingOID},
			bson.M{"$set": bson.M{"lastModified": time.Now(), "isFriend": false}},
		)
		if err != nil {
			return res1, err
		}
		res2, err := mongoDB.Followings.DeleteOne(
			ctx,
			bson.M{"userID": userOID, "followingID": followingOID},
		)
		if err != nil {
			return res2, err
		}
		res3, err := mongoDB.Users.UpdateOne(
			ctx,
			bson.M{"_id": userOID},
			bson.M{"$inc": bson.M{"followings": -1}},
		)
		if err != nil {
			return res3, err
		}
		res4, err := mongoDB.Users.UpdateOne(
			ctx,
			bson.M{"_id": followingOID},
			bson.M{"$inc": bson.M{"followers": -1}},
		)
		if err != nil {
			return res4, err
		}
		if res2.DeletedCount != 1 || res3.ModifiedCount == 1 || res4.ModifiedCount != 1 {
			return nil, errors.New(
				strings.Join(
					[]string{
						"delete:",
						strconv.FormatInt(res2.DeletedCount, 10),
						"inc followings:",
						strconv.FormatInt(res3.ModifiedCount, 10),
						"inc followers:",
						strconv.FormatInt(res4.ModifiedCount, 10),
					},
					" ",
				),
			)
		}
		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	return err
}

func GetFriends(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Followings.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.M{"userID": oID, "isFriend": true}}},
			bson.D{{Key: "$sort", Value: bson.M{"lastModified": -1}}},
			bson.D{{Key: "$skip", Value: index}},
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "users",
					"localField":   "followingID",
					"foreignField": "_id",
					"as":           "user",
				},
			}},
			bson.D{{Key: "$replaceRoot", Value: bson.M{"newRoot": bson.M{"$first": "$user"}}}},
		},
		options.Aggregate().SetMaxTime(2*time.Second),
	)
}
