package model

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	Users     *mongo.Collection
	Instants  *mongo.Collection
	Comments  *mongo.Collection
	Following *mongo.Collection
}
