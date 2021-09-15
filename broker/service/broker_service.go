package service

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/database/mgdb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BrokerService interface {
	InsertOne(worker *model.Broker) (*mongo.InsertOneResult, error)
	/*
	*@filter primitive.ObjectIDFromHex("60acb63ad1b5adedd2da8766")
	 */
	FindOne(filter bson.M) (*model.Broker, error)
	FindAll(filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error)
}

type Broker struct {
}

func (s *Broker) InsertOne(broker *model.Broker) (*mongo.InsertOneResult, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Database: mgdb.Database}
	return brokerRepo.InsertOne(context.TODO(), broker)
}

func (s *Broker) FindOne(filter bson.M) (*model.Broker, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Database: mgdb.Database}
	return brokerRepo.FindOne(context.TODO(), filter)
}

func (s *Broker) FindAll(filter bson.M, page *model.PageableRequest) (*model.BrokerItemsPage, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Database: mgdb.Database}
	return brokerRepo.FindAll(context.TODO(), filter, page)
}

func (s *Broker) Register() {
	// Query Database
	filter := bson.M{"broker_link": config.GetGrpcLink(), "broker_label": config.GetBrokerLabel()}
	r, _ := s.FindOne(filter)
	if r != nil {
		config.SetBrokerId(r.BrokerId)
	} else {
		config.SetBrokerId(uuid.New().String())
		// Save worker's information to DB
		broker := &model.Broker{
			BrokerId:    config.GetBrokerId(),
			BrokerLink:  config.GetGrpcLink(),
			BrokerLabel: config.GetBrokerLabel(),
			CreateAt:    ptypes.TimestampNow(),
			UpdateAt:    ptypes.TimestampNow(),
		}
		_, _ = s.InsertOne(broker)
	}
}
