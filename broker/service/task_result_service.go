package service

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskResultService interface {
	InsertOne(worker *model.TaskResult) (*mongo.InsertOneResult, error)
}

type TaskResult struct {
}

func (s *TaskResult) InsertOne(tr *model.TaskResult) (*mongo.InsertOneResult, error) {
	var taskResultRepo repository.TaskResultRepository = &repository.TaskResultMongo{Client: mongodb.Client}
	return taskResultRepo.InsertOne(context.TODO(), tr)
}
func (s *TaskResult) FindAll(filter bson.M, page *model.PageableRequest) (*model.TaskResultItemsPage, error) {
	var taskResultRepo repository.TaskResultRepository = &repository.TaskResultMongo{Client: mongodb.Client}
	return taskResultRepo.FindAll(context.TODO(), filter, page)
}
