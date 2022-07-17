package database

import (
	"time"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetComments(insID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	return mongoDB.Comments.Find(
		ctx,
		bson.M{"insID": insID},
		options.Find().SetSort(bson.M{"_id": -1}).SetSkip(index).SetLimit(pageSize),
	)
}

func PostComment(comment model.Comment) (*mongo.InsertOneResult, error) {
	instantOID, err := primitive.ObjectIDFromHex(comment.InsID)
	if err != nil {
		return nil, err
	}
	userOID, err := primitive.ObjectIDFromHex(comment.UserID)
	if err != nil {
		return nil, err
	}
	replyToOID, err := primitive.ObjectIDFromHex(comment.ReplyToID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return mongoDB.Comments.InsertOne(
		ctx,
		bson.M{
			"created":      now,
			"lastModified": now,
			"insID":        instantOID,
			"userID":       userOID,
			"content":      comment.Content,
			"replyToID":    replyToOID,
			"direct":       comment.Direct,
		},
	)
}
