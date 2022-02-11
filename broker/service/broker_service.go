package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/database/mgdb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BrokerService interface {
	InsertOne(worker *model.Broker) (*mongo.InsertOneResult, error)
	UpdateOne(filter bson.M, data bson.M) (*mongo.UpdateResult, error)
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

func (s *Broker) UpdateOne(filter bson.M, data bson.M) (*mongo.UpdateResult, error) {
	var brokerRepo repository.BrokerRepository = &repository.BrokerMongo{Database: mgdb.Database}
	return brokerRepo.UpdateOne(context.TODO(), filter, data)
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
	filter := bson.M{"broker_id": config.GetBrokerId()}
	r, _ := s.FindOne(filter)
	if r != nil {
		filter1 := bson.M{"broker_id": r.BrokerId}
		update := bson.M{"$set": bson.M{"status": 1, "broker_link": config.GetGrpcLink(), "broker_label": config.GetBrokerLabel(), "update_at": timestamppb.Now(), "version": config.GetVersion()}}
		_, _ = s.UpdateOne(filter1, update)
	} else {
		config.SetBrokerId(uuid.New().String())
		// Save worker's information to DB
		broker := &model.Broker{
			BrokerId:    config.GetBrokerId(),
			BrokerLink:  config.GetGrpcLink(),
			BrokerLabel: config.GetBrokerLabel(),
			Version:     config.GetVersion(),
			Status:      1,
			CreateAt:    timestamppb.Now(),
			UpdateAt:    timestamppb.Now(),
		}
		_, _ = s.InsertOne(broker)
	}
}
