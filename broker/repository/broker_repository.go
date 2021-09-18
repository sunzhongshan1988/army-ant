package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type BrokerRepository interface {
	InsertOne(ctx context.Context, worker *model.Broker) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter bson.M) (*model.Broker, error)
	FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error)
}

type BrokerMongo struct {
	Database *mongo.Database
}

func (r *BrokerMongo) InsertOne(ctx context.Context, broker *model.Broker) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Database.Collection("broker").InsertOne(ctx, broker)
	if err != nil {
		log.Printf("[mongodb, save] error:%v", err)
	}

	log.Printf("[mongodb, save] info: %v", insertResult.InsertedID)

	return insertResult, err
}

func (r *BrokerMongo) FindOne(ctx context.Context, filter bson.M) (*model.Broker, error) {

	var result model.Broker

	err := r.Database.Collection("broker").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[mongodb, findone] error: %v", err)
	}

	log.Printf("[mongodb, findone] info: %v", "success")

	return &result, err
}

func (r *BrokerMongo) FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error) {

	result := model.BrokerItemsPage{}

	findOptions := &options.FindOptions{}
	findOptions.SetSort(bson.M{"_id": -1})
	if page.Size > 0 {
		findOptions.SetLimit(page.Size)
		findOptions.SetSkip(page.Index * page.Size)
	}

	cur, err := r.Database.Collection("broker").Find(context.TODO(), filter, findOptions)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[mongodb, finall] error:%v", err)
	}

	count, err1 := r.Database.Collection("broker").CountDocuments(ctx, filter)
	if err1 != nil {
		log.Printf("[mongodb, finall] error:%v", err)
	}

	result.Total = count
	result.PageSize = page.Size
	result.Current = page.Index

	defer cur.Close(ctx)
	if err = cur.All(ctx, &result.Items); err != nil {
		log.Printf("[mongodb, finall] error:%v", err)
	}

	log.Printf("[mongodb, finall] info: %v", "success")

	return &result, err
}
