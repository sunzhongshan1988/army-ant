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
	FindById(ctx context.Context, filter bson.M)
}

type WorkerRepository struct {
	Client *mongo.Client
}

func (r *WorkerRepository) InsertOne(ctx context.Context, worker *model.WorkerRegister) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Client.Database("armyant").Collection("worker").InsertOne(ctx, worker)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("MongoDB Save: %v", insertResult.InsertedID)

	return insertResult, nil
}

func (r *WorkerRepository) FindById(ctx context.Context, filter bson.M) (*model.WorkerRegister, error) {

	var result model.WorkerRegister

	err := r.Client.Database("armyant").Collection("worker").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Printf("MongoDB FindOne: record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	log.Printf("MongoDB FindOne: %v", "success")

	return &result, nil
}
