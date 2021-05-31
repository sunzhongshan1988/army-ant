package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type taskResultRepository interface {
	InsertOne(ctx context.Context, tr *model.TaskResult) (*mongo.InsertOneResult, error)
}

type TaskResultRepository struct {
	Client *mongo.Client
}

func (r *TaskResultRepository) InsertOne(ctx context.Context, worker *model.TaskResult) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Client.Database("armyant").Collection("task_result").InsertOne(ctx, worker)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("MongoDB Save: %v", insertResult.InsertedID)

	return insertResult, nil
}
