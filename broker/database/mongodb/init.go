package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var client *mongo.Client
var collection *mongo.Collection
var ctx = context.TODO()

func Init() {
	var cancel context.CancelFunc
	var err error

	uri := "mongodb://armyant:%40WSX3edc@10.11.51.152:27017/armyant?authSource=admin&readPreference=primary&appname=ArmyAnt&ssl=false"

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Printf("MongoDB: Successfully connected and pinged.")
}
