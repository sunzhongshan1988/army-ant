package service

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/database/mgdb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkerService interface {
	InsertOne(worker *model.Worker) (*mongo.InsertOneResult, error)
	/*
	*@filter primitive.ObjectIDFromHex("60acb63ad1b5adedd2da8766")
	 */
	FindOne(filter bson.M) (*model.Worker, error)
}

type Worker struct {
}

func (s *Worker) InsertOne(worker *model.Worker) (*mongo.InsertOneResult, error) {
	var workerRepo repository.WorkerRepository = &repository.WorkerMongo{Client: mgdb.Client}
	return workerRepo.InsertOne(context.TODO(), worker)
}

func (s *Worker) FindOne(filter bson.M) (*model.Worker, error) {
	var workerRepo repository.WorkerRepository = &repository.WorkerMongo{Client: mgdb.Client}
	return workerRepo.FindOne(context.TODO(), filter)
}

func (s *Worker) FindAll(filter bson.M, page *model.PageableRequest) (*model.WorkerItemsPage, error) {
	var workerRepo repository.WorkerRepository = &repository.WorkerMongo{Client: mgdb.Client}
	return workerRepo.FindAll(context.TODO(), filter, page)
}
