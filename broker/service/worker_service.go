package service

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type workerService interface {
	InsertOne(worker *model.Worker) (*mongo.InsertOneResult, error)
	/*
	*@filter primitive.ObjectIDFromHex("60acb63ad1b5adedd2da8766")
	 */
	FindOne(filter bson.M) (*model.Worker, error)
}

type WorkerService struct {
}

func (s *WorkerService) InsertOne(worker *model.Worker) (*mongo.InsertOneResult, error) {
	var workerRepo = repository.WorkerRepository{Client: mongodb.Client}
	return workerRepo.InsertOne(context.TODO(), worker)
}

func (s *WorkerService) FindOne(filter bson.M) (*model.Worker, error) {
	var workerRepo = repository.WorkerRepository{Client: mongodb.Client}
	return workerRepo.FindOne(context.TODO(), filter)
}
