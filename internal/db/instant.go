package database

import (
	"time"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetInstants(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Instants.Find(ctx, bson.M{"userID": userID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func PostInstant(instant model.Instant) (*mongo.InsertOneResult, error) {
	return mongoDB.Instants.InsertOne(ctx, bson.M{"userID": instant.UserID, "created": time.Now(), "content": instant.Content})
}

func UpdateInstant(instant model.Instant) (*mongo.UpdateResult, error) {
	return mongoDB.Instants.UpdateOne(ctx, bson.M{"_id":instant.InsID }, bson.M{"$set": bson.M{"content": instant.Content}, "$currentDate": bson.M{"lastModified": true}})
}

func LikeInstant(like model.Like) (*mongo.UpdateResult, error) {
	return mongoDB.Instants.UpdateOne(ctx, bson.M{"insID": like.InsID}, bson.M{"$set": bson.M{"useID": like.UserID, "attitude": like.Attitude}, "$currentDate": bson.M{"lastModified": true}})
}

func ShareInstant(instant model.Instant) (*mongo.InsertOneResult, error) {
	return mongoDB.Instants.InsertOne(ctx, bson.M{"userID": instant.UserID, "created": time.Now(), "content": instant.Content, "refOriginID":instant.RefOriginID})
}
