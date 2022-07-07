package model

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	Client     *mongo.Client
	Users      *mongo.Collection
	Instants   *mongo.Collection
	Comments   *mongo.Collection
	Followings *mongo.Collection
	Likes      *mongo.Collection
}
