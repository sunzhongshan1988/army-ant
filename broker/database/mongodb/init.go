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
var Ctx = context.TODO()

func Init() {
	var cancel context.CancelFunc
	var err error

	uri := "mongodb://armyant:%40WSX3edc@10.11.51.152:27017/armyant?authSource=admin&readPreference=primary&appname=ArmyAnt&ssl=false"

	Ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(Ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(Ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Printf("MongoDB: Successfully connected and pinged.")
}

func GetCollection(collectionName string) (collection *mongo.Collection) {
	collection = client.Database("armyant").Collection(collectionName)
	return collection
}
