package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type workerRepository interface {
	InsertOne(ctx context.Context, worker *model.WorkerRegister) (*mongo.InsertOneResult, error)
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
