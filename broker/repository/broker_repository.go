package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
)

type BrokerRepository interface {
	InsertOne(ctx context.Context, worker *model.Broker) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter bson.M) (*model.Broker, error)
	FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error)
}

type BrokerMongo struct {
	Client *mongo.Client
}

func (r *BrokerMongo) InsertOne(ctx context.Context, broker *model.Broker) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Client.Database("armyant").Collection("broker").InsertOne(ctx, broker)
	if err != nil {
		log.Printf("[mongodb,save] error:%v", err)
	}

	log.Printf("[mongodb,save] info: %v", insertResult.InsertedID)

	return insertResult, nil
}

func (r *BrokerMongo) FindOne(ctx context.Context, filter bson.M) (*model.Broker, error) {

	var result model.Broker

	err := r.Client.Database("armyant").Collection("broker").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,db]%v", err)
	}

	log.Printf("MongoDB FindOne: %v", "success")

	return &result, nil
}

func (r *BrokerMongo) FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error) {

	result := model.BrokerItemsPage{}

	findOptions := &options.FindOptions{}
	if page.Size > 0 {
		findOptions.SetLimit(page.Size)
		findOptions.SetSkip(page.Index * page.Size)
	}

	cur, err := r.Client.Database("armyant").Collection("broker").Find(context.TODO(), filter, findOptions)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,db]%v", err)
	}

	count, err1 := r.Client.Database("armyant").Collection("broker").CountDocuments(ctx, filter)
	if err1 != nil {
		log.Printf("[error,db]%v", err)
	}

	result.TotalItems = count
	result.TotalPages = int64(math.Ceil(float64(count) / float64(page.Size)))
	result.CurrentPage = page.Index

	defer cur.Close(ctx)
	if err = cur.All(ctx, &result.Items); err != nil {
		log.Printf("[error,db]%v", err)
	}

	log.Printf("MongoDB Find: %v", "success")

	return &result, nil
}
