package config

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"go.mongodb.org/mongo-driver/bson"
)

type Broker struct {
	label       string
	brokerId    string
	brokerType  string
	address     string
	grpcPort    string
	graphqlPort string
}

var broker = &Broker{}

func Init() {
	broker.label = "system"
	broker.brokerId = ""
	broker.brokerType = "main"
	broker.address = "127.0.0.1"
	broker.grpcPort = "50051"
	broker.graphqlPort = "8080"

	registerBroker()
}

func GetBrokerLabel() string {
	return broker.label
}
func SetBrokerLabel(lable string) {
	broker.label = lable
}

func GetBrokerId() string {
	return broker.brokerId
}
func SetBrokerId(id string) {
	broker.brokerId = id
}

func GetBrokerType() string {
	return broker.brokerType
}

func GetAddress() string {
	return broker.address
}

func GetGrpcPort() string {
	return broker.grpcPort
}

func GetGrpcLink() string {
	return broker.address + ":" + broker.grpcPort
}

func GetGraphQLPort() string {
	return broker.graphqlPort
}

func registerBroker() {
	var brokerService = service.BrokerService{}
	broker := &model.BrokerRegister{
		BrokerId:    GetBrokerId(),
		BrokerLink:  GetGrpcLink(),
		BrokerLabel: GetBrokerLabel(),
		CreateAt:    ptypes.TimestampNow(),
		UpdateAt:    ptypes.TimestampNow(),
	}

	// Query Database
	filter := bson.M{"broker_link": GetGrpcLink(), "broker_label": GetBrokerLabel()}
	r, _ := brokerService.FindOne(filter)
	if r != nil {
		SetBrokerId(r.BrokerId)
	} else {
		SetBrokerId(uuid.New().String())
		// Save worker's information to DB
		_, _ = brokerService.InsertOne(broker)
	}
}
