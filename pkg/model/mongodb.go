package model

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	Users     *mongo.Collection
	Instants  *mongo.Collection
	Comments  *mongo.Collection
	Followings *mongo.Collection
	Likes *mongo.Collection
}
