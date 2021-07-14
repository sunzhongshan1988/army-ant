package config

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
)

type Broker struct {
	Label       string `json:"label"`
	BrokerId    string `json:"brokerId"`
	BrokerType  string `json:"brokerType"`
	Address     string `json:"address"`
	GrpcPort    string `json:"grpcPort"`
	GraphqlPort string `json:"graphqlPort"`
}

var broker = &Broker{}

func Init() {
	broker.Label = os.Getenv("AAB_LABEL") // "system"
	broker.BrokerId = ""
	broker.BrokerType = "main"
	broker.Address = os.Getenv("AAB_ADDRESS")          // "127.0.0.1"
	broker.GrpcPort = os.Getenv("AAB_GRPC_PORT")       // "50051"
	broker.GraphqlPort = os.Getenv("AAB_GRAPHQL_PORT") // "8080"

	registerBroker()

	jsonStr, _ := json.Marshal(broker)
	log.Printf("[config, init] info: %v", string(jsonStr))
}

func GetBrokerLabel() string {
	return broker.Label
}
func SetBrokerLabel(lable string) {
	broker.Label = lable
}

func GetBrokerId() string {
	return broker.BrokerId
}
func SetBrokerId(id string) {
	broker.BrokerId = id
}

func GetBrokerType() string {
	return broker.BrokerType
}

func GetAddress() string {
	return broker.Address
}

func GetGrpcPort() string {
	return broker.GrpcPort
}

func GetGrpcLink() string {
	return broker.Address + ":" + broker.GrpcPort
}

func GetGraphQLPort() string {
	return broker.GraphqlPort
}

func registerBroker() {
	brokerService := service.Broker{}

	// Query Database
	filter := bson.M{"broker_link": GetGrpcLink(), "broker_label": GetBrokerLabel()}
	r, _ := brokerService.FindOne(filter)
	if r != nil {
		SetBrokerId(r.BrokerId)
	} else {
		SetBrokerId(uuid.New().String())
		// Save worker's information to DB
		broker := &model.Broker{
			BrokerId:    GetBrokerId(),
			BrokerLink:  GetGrpcLink(),
			BrokerLabel: GetBrokerLabel(),
			CreateAt:    ptypes.TimestampNow(),
			UpdateAt:    ptypes.TimestampNow(),
		}
		_, _ = brokerService.InsertOne(broker)
	}
}
