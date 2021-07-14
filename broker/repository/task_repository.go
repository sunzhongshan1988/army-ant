package repository

import (
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
)

type TaskRepository interface {
	FindOne(ctx context.Context, filter bson.M) (*model.Task, error)
	InsertOne(ctx context.Context, tr *model.Task) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error)
}

type TaskMongo struct {
	Client *mongo.Client
}

func (r *TaskMongo) FindOne(ctx context.Context, filter bson.M) (*model.Task, error) {

	var result model.Task

	err := r.Client.Database("armyant").Collection("task").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[mongodb, findone] error: %v", err)
	}

	log.Printf("[mongodb, findone] info: %v", "success")

	return &result, err
}

func (r *TaskMongo) InsertOne(ctx context.Context, worker *model.Task) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Client.Database("armyant").Collection("task").InsertOne(ctx, worker)

	if err != nil {
		log.Printf("[mongodb,save] error:%v", err)
	}

	log.Printf("[mongodb,save] info: %v", insertResult.InsertedID)

	return insertResult, err
}

func (r *TaskMongo) UpdateOne(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error) {

	updateResult, err := r.Client.Database("armyant").Collection("task").UpdateOne(ctx, filter, data)

	if err != nil {
		log.Printf("[mongodb,updateone] error:%v", err)
	}

	log.Printf("[mongodb,updateone] info: %v", updateResult.ModifiedCount)

	return updateResult, err
}
