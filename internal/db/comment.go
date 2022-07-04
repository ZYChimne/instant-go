package database

import (
	"time"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetComments(insID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Comments.Find(ctx, bson.M{"insID": insID}, options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize))
}

func PostComment(comment model.Comment) (*mongo.InsertOneResult, error) {
	return mongoDB.Comments.InsertOne(ctx, bson.M{"created":time.Now(), "insID":comment.InsID, "userID":comment.UserID, "content":comment.Content})
}
