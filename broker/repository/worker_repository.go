package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type WorkerRepository interface {
	InsertOne(ctx context.Context, worker *model.WorkerRegister) (*mongo.InsertOneResult, error)
}

type workerRepository struct {
	client *mongo.Client
}

func (r *workerRepository) InsertOne(ctx context.Context, worker *model.WorkerRegister) (*mongo.InsertOneResult, error) {

	insertResult, err := r.client.Database("armyant").Collection("worker").InsertOne(ctx, worker)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("MongoDB Save: %v", insertResult.InsertedID)

	return insertResult, nil
}
