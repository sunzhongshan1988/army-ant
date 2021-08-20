package mgdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Client *mongo.Client

var Database = "armyant_dev"

func Init() *mongo.Client {

	uri := "mongodb://armyant:%40WSX3edc@10.11.51.152:27017/" + Database + "?authSource=admin&readPreference=primary&appname=ArmyAnt&ssl=false"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err = Client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	// Ping the primary
	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Printf("[mgdb, init] info: Successfully connected and pinged.")

	return Client
}
