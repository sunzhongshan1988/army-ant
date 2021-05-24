package service

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkerService struct {
}

var repo = repository.WorkerRepository{Client: mongodb.GetClient()}

func (s *WorkerService) InsertOne(worker *model.WorkerRegister) (*mongo.InsertOneResult, error) {
	return repo.InsertOne(context.TODO(), worker)
}
