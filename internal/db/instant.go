package database

import (
	"errors"
	"log"
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
	return mongoDB.Feeds.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{
				{Key: "$match", Value: bson.M{"userID": oID}},
			},
			bson.D{{Key: "$unwind", Value: "$instants"}},
			bson.D{{Key: "$sort", Value: bson.M{"instants.insID": -1}}},
			bson.D{{Key: "$skip", Value: index}},
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "instants",
					"localField":   "instants.insID",
					"foreignField": "_id",
					"as":           "feeds",
				},
			}},
			bson.D{
				{
					Key: "$replaceRoot",
					Value: bson.M{
						"newRoot": bson.M{
							"$mergeObjects": bson.A{bson.M{"$first": "$feeds"}, "$instants"},
						},
					},
				},
			}},
		options.Aggregate().SetMaxTime(time.Second*2),
	)
}

func GetInstantsByUserID(userID string, index int64, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Instants.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{
				{Key: "$match", Value: bson.M{"userID": oID}},
			},
			bson.D{{Key: "$sort", Value: bson.M{"instants.insID": -1}}},
			bson.D{{Key: "$skip", Value: index}},
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "likes",
					"localField":   "_id",
					"foreignField": "insID",
					"as":           "likeList",
					"pipeline": bson.A{
						bson.D{
							{Key: "$match", Value: bson.M{"userID": oID}},
						},
					},
				},
			}},
			bson.D{
				{
					Key: "$replaceRoot",
					Value: bson.M{
						"newRoot": bson.M{
							"$mergeObjects": bson.A{bson.M{"$first": "$likeList"}, "$$ROOT"},
						},
					},
				},
			}},
		options.Aggregate().SetMaxTime(time.Second*2),
	)
}

func PostInstant(instant model.Instant) error {
	oID, err := primitive.ObjectIDFromHex(instant.UserID)
	if err != nil {
		return err
	}
	var user model.User
	err = mongoDB.Users.FindOne(ctx, bson.M{"_id": oID}).Decode(&user)
	if err != nil {
		return err
	}
	session, err := mongoDB.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		res1, err := mongoDB.Instants.InsertOne(
			ctx,
			bson.M{
				"userID":       oID,
				"username":     user.Username,
				"avatar":       user.Avatar,
				"created":      time.Now(),
				"lastModified": time.Now(),
				"content":      instant.Content,
				"likes":        0,
				"shares":       0,
			},
		)
		if err != nil {
			return res1, nil
		}
		rows, err := mongoDB.Followings.Find(
			ctx,
			bson.M{"followingID": oID},
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close(ctx)
		for rows.Next(ctx) {
			err = fanOutOnWrite(
				res1.InsertedID.(primitive.ObjectID),
				rows.Current.Lookup("userID").ObjectID(),
			)
			if err != nil {
				return nil, err
			}
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		err = fanOutOnWrite(res1.InsertedID.(primitive.ObjectID), oID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	return err
}

func fanOutOnWrite(instantOID primitive.ObjectID, userOID primitive.ObjectID) error {
	rows, err := mongoDB.Feeds.Aggregate(ctx, mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.M{"userID": userOID}},
		},
		bson.D{{Key: "$limit", Value: 1}},
		bson.D{
			{
				Key: "$project",
				Value: bson.M{
					"size": bson.M{
						"$cond": bson.M{
							"if":   bson.M{"$isArray": "$instants"},
							"then": bson.M{"$size": "$instants"},
							"else": 0,
						},
					},
				},
			},
		}},
		options.Aggregate().SetMaxTime(time.Second*2))
	if err != nil {
		return err
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		size := rows.Current.Lookup("size").Int32()
		log.Println(size)
		if size >= maxFeedSize {
			_, err := mongoDB.Feeds.UpdateOne(
				ctx,
				bson.M{"userID": userOID},
				bson.M{
					"$pop": bson.M{"instants": -1},
				},
				options.Update().SetUpsert(true),
			)
			if err != nil {
				return err
			}
		}
	}
	_, err = mongoDB.Feeds.UpdateOne(
		ctx,
		bson.M{"userID": userOID},
		bson.M{
			"$push": bson.M{"instants": bson.M{"insID": instantOID, "attitude": 0}},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	return nil
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
	return mongoDB.Instants.UpdateOne(
		ctx,
		bson.M{"_id": instantOID, "userID": userOID},
		bson.M{
			"$set":         bson.M{"content": instant.Content},
			"$currentDate": bson.M{"lastModified": true},
		},
	)
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
		res1, err := mongoDB.Likes.UpdateOne(
			ctx,
			bson.M{"insID": instantOID},
			bson.M{
				"$set":         bson.M{"userID": userOID, "attitude": like.Attitude},
				"$currentDate": bson.M{"lastModified": true},
			},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return res1, nil
		}
		res2, err := mongoDB.Instants.UpdateOne(
			ctx,
			bson.M{"_id": instantOID},
			bson.M{"$inc": bson.M{"likes": 1}},
		)
		if err != nil {
			return res2, nil
		}
		res3, err := mongoDB.Feeds.UpdateOne(
			ctx,
			bson.M{"userID": userOID, "instants.insID": instantOID},
			bson.M{"$set": bson.M{"instants.$.attitude": like.Attitude}},
		)
		if err != nil {
			return res3, err
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
	return mongoDB.Instants.InsertOne(
		ctx,
		bson.M{
			"userID":      userOID,
			"created":     time.Now(),
			"content":     instant.Content,
			"refOriginID": instantOID,
		},
	)
}

func GetLikesUsername(instantID string, pageSize int64) (*mongo.Cursor, error) {
	oID, err := primitive.ObjectIDFromHex(instantID)
	if err != nil {
		return nil, err
	}
	return mongoDB.Likes.Aggregate(
		ctx,
		mongo.Pipeline{
			bson.D{
				{Key: "$match", Value: bson.M{"insID": oID}},
			},
			bson.D{{Key: "$sort", Value: bson.M{"_id": -1}}},
			bson.D{{Key: "$limit", Value: pageSize}},
			bson.D{{
				Key: "$lookup",
				Value: bson.M{
					"from":         "users",
					"localField":   "userID",
					"foreignField": "_id",
					"as":           "users",
				},
			}},
			bson.D{
				{
					Key: "$replaceRoot",
					Value: bson.M{
						"newRoot": bson.M{
							"$first": "$users",
						},
					},
				},
			},
			bson.D{
				{
					Key: "$project",
					Value: bson.M{
						"_id":      0,
						"username": 1,
					},
				},
			}},
		options.Aggregate().SetMaxTime(time.Second*2),
	)
}
