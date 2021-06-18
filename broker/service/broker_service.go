package service

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type brokerService interface {
	InsertOne(worker *model.BrokerRegister) (*mongo.InsertOneResult, error)
	/*
	*@filter primitive.ObjectIDFromHex("60acb63ad1b5adedd2da8766")
	 */
	FindOne(filter bson.M) (*model.BrokerRegister, error)
}

type BrokerService struct {
}

func (s *BrokerService) InsertOne(broker *model.BrokerRegister) (*mongo.InsertOneResult, error) {
	var workerRepo = repository.BrokerRepository{Client: mongodb.Client}
	return workerRepo.InsertOne(context.TODO(), broker)
}

func (s *BrokerService) FindOne(filter bson.M) (*model.BrokerRegister, error) {
	var workerRepo = repository.BrokerRepository{Client: mongodb.Client}
	return workerRepo.FindOne(context.TODO(), filter)
}
