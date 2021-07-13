package service

import (
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type TaskService interface {
	FindOne(filter bson.M) (*model.Task, error)
	InsertOne(worker *model.Task) (*mongo.InsertOneResult, error)
	UpdateOne(filter bson.M, data bson.M) (*mongo.UpdateResult, error)
}

type Task struct {
}

func (s *Task) FindOne(filter bson.M) (*model.Task, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mongodb.Client}
	return taskRepo.FindOne(context.TODO(), filter)
}

func (s *Task) InsertOne(tr *model.Task) (*mongo.InsertOneResult, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mongodb.Client}
	return taskRepo.InsertOne(context.TODO(), tr)
}
func (s *Task) UpdateOne(filter bson.M, data bson.M) (*mongo.UpdateResult, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mongodb.Client}
	return taskRepo.UpdateOne(context.TODO(), filter, data)
}
