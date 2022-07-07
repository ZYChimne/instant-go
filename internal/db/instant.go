package database

import (
	"errors"
	"strings"
	"time"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetInstants(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Instants.Find(ctx, bson.M{"userID": oID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func PostInstant(instant model.Instant) (*mongo.InsertOneResult, error) {
	oID, err := primitive.ObjectIDFromHex(instant.UserID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Instants.InsertOne(ctx, bson.M{"userID": oID, "created": time.Now(), "lastModified": time.Now(), "content": instant.Content, "likes": 0, "shares": 0})
}

func UpdateInstant(instant model.Instant) (*mongo.UpdateResult, error) {
	userOID, err := primitive.ObjectIDFromHex(instant.UserID)
	if err != nil {
		return nil, err
	}
	instantOID, err := primitive.ObjectIDFromHex(instant.InsID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Instants.UpdateOne(ctx, bson.M{"_id": instantOID, "userID": userOID}, bson.M{"$set": bson.M{"content": instant.Content}, "$currentDate": bson.M{"lastModified": true}})
}

func LikeInstant(like model.Like) error {
	session, err := mongoDB.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		userOID, err := primitive.ObjectIDFromHex(like.UserID)
		if err != nil {
			return nil, err
		}
		instantOID, err := primitive.ObjectIDFromHex(like.InsID)
		if err != nil {
			return nil, err
		}
		res1, err := mongoDB.Likes.UpdateOne(ctx, bson.M{"insID": instantOID}, bson.M{"$set": bson.M{"useID": userOID, "attitude": like.Attitude}, "$currentDate": bson.M{"lastModified": true}}, options.Update().SetUpsert(true))
		if err != nil {
			return res1, nil
		}
		res2, err := mongoDB.Instants.UpdateOne(ctx, bson.M{"_id": instantOID}, bson.M{"$inc": bson.M{"likes": 1}})
		if err != nil {
			return res2, nil
		}
		if (res1.UpsertedCount+res1.ModifiedCount != 1) || res2.ModifiedCount == 0 {
			return nil, errors.New(strings.Join([]string{"", "instant not found"}, " "))
		}
		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	return err
}

func ShareInstant(instant model.Instant) (*mongo.InsertOneResult, error) {
	userOID, err := primitive.ObjectIDFromHex(instant.UserID)
	if err != nil {
		return nil, err
	}
	instantOID, err := primitive.ObjectIDFromHex(instant.RefOriginID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Instants.InsertOne(ctx, bson.M{"userID": userOID, "created": time.Now(), "content": instant.Content, "refOriginID": instantOID})
}
