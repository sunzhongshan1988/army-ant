package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type workerRepository interface {
	InsertOne(ctx context.Context, worker *model.WorkerRegister) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter bson.M)
}

type WorkerRepository struct {
	Client *mongo.Client
}

func (r *WorkerRepository) InsertOne(ctx context.Context, worker *model.WorkerRegister) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Client.Database("armyant").Collection("worker").InsertOne(ctx, worker)
	if err != nil {
		log.Printf("[error,db]%v", err)
	}

	log.Printf("MongoDB Save: %v", insertResult.InsertedID)

	return insertResult, nil
}

func (r *WorkerRepository) FindOne(ctx context.Context, filter bson.M) (*model.WorkerRegister, error) {

	var result model.WorkerRegister

	err := r.Client.Database("armyant").Collection("worker").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,db]%v", err)
	}

	log.Printf("MongoDB FindOne: %v", "success")

	return &result, nil
}
