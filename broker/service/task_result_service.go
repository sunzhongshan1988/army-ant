package service

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskResultService interface {
	InsertOne(worker *model.TaskResult) (*mongo.InsertOneResult, error)
}

type TaskResultService struct {
}

func (s *TaskResultService) InsertOne(tr *model.TaskResult) (*mongo.InsertOneResult, error) {
	var taskResultRepo = repository.TaskResultRepository{Client: mongodb.Client}
	return taskResultRepo.InsertOne(context.TODO(), tr)
}
