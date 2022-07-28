package database

import (
	"context"
	"log"
	"time"
	"zychimne/instant/pkg/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *model.MongoDB

func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Connect MongoDB error ", err)
		return
	}
	database := client.Database("instant")
	mongoDB = &model.MongoDB{
		Client:     client,
		Users:      database.Collection("users"),
		Instants:   database.Collection("instants"),
		Feeds:      database.Collection("feeds"),
		Sharings:   database.Collection("sharings"),
		Comments:   database.Collection("comments"),
		Followings: database.Collection("followings"),
		Likes:      database.Collection("likes"),
	}
}
