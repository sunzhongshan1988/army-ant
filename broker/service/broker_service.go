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
	InsertOne(worker *model.Broker) (*mongo.InsertOneResult, error)
	/*
	*@filter primitive.ObjectIDFromHex("60acb63ad1b5adedd2da8766")
	 */
	FindOne(filter bson.M) (*model.Broker, error)
	FindAll(filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error)
}

type BrokerService struct {
}

func (s *BrokerService) InsertOne(broker *model.Broker) (*mongo.InsertOneResult, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Client: mongodb.Client}
	return brokerRepo.InsertOne(context.TODO(), broker)
}

func (s *BrokerService) FindOne(filter bson.M) (*model.Broker, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Client: mongodb.Client}
	return brokerRepo.FindOne(context.TODO(), filter)
}

func (s *BrokerService) FindAll(filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Client: mongodb.Client}
	return brokerRepo.FindAll(context.TODO(), filter, page)
}
