package database

import (
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
	return mongoDB.Instants.InsertOne(ctx, bson.M{"userID": oID, "created": time.Now(), "lastModified": time.Now(), "content": instant.Content})
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
	return mongoDB.Instants.UpdateOne(ctx, bson.M{"_id": instantOID, "userID":userOID}, bson.M{"$set": bson.M{"content": instant.Content}, "$currentDate": bson.M{"lastModified": true}})
}

func LikeInstant(like model.Like) (*mongo.UpdateResult, error) {
	userOID, err := primitive.ObjectIDFromHex(like.UserID)
	if err != nil {
		return nil, err
	}
	instantOID, err := primitive.ObjectIDFromHex(like.InsID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Instants.UpdateOne(ctx, bson.M{"insID": instantOID}, bson.M{"$set": bson.M{"useID": userOID, "attitude": like.Attitude}, "$currentDate": bson.M{"lastModified": true}})
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
