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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("instant")
	mongoDB = &model.MongoDB{Users: database.Collection("users"), Instants: database.Collection("instants")}
}
