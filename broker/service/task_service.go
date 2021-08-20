package service

import (
	"github.com/sunzhongshan1988/army-ant/broker/database/mgdb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type TaskService interface {
	AnalyseTaskStatus(pipeline mongo.Pipeline) (*[]model.AnalyseTaskStatus, error)
	FindOne(filter bson.M) (*model.Task, error)
	InsertOne(worker *model.Task) (*mongo.InsertOneResult, error)
	UpdateOne(filter bson.M, data bson.M) (*mongo.UpdateResult, error)
}

type Task struct {
}

func (s *Task) AnalyseTaskStatus(pipeline mongo.Pipeline) (*[]model.AnalyseTaskStatus, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mgdb.Client}
	return taskRepo.AnalyseTaskStatus(context.TODO(), pipeline)
}

func (s *Task) FindOne(filter bson.M) (*model.Task, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mgdb.Client}
	return taskRepo.FindOne(context.TODO(), filter)
}

func (s *Task) FindAll(filter bson.M, page *model.PageableRequest) (*model.TaskItemsPage, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mgdb.Client}
	return taskRepo.FindAll(context.TODO(), filter, page)
}

func (s *Task) InsertOne(tr *model.Task) (*mongo.InsertOneResult, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mgdb.Client}
	return taskRepo.InsertOne(context.TODO(), tr)
}
func (s *Task) UpdateOne(filter bson.M, data bson.M) (*mongo.UpdateResult, error) {
	var taskRepo repository.TaskRepository = &repository.TaskMongo{Client: mgdb.Client}
	return taskRepo.UpdateOne(context.TODO(), filter, data)
}
