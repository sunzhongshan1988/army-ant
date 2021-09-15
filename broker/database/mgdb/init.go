package mgdb

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Client *mongo.Client
var Database *mongo.Database

func Init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongodbUri()))
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

	Database = Client.Database(config.GetMongodbDatabase())
}
