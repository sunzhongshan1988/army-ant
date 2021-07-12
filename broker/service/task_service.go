package service

import (
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type TaskService interface {
	InsertOne(worker *model.Task) (*mongo.InsertOneResult, error)
}

type Task struct {
}

func (s *Task) InsertOne(tr *model.Task) (*mongo.InsertOneResult, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mongodb.Client}
	return taskRepo.InsertOne(context.TODO(), tr)
}
