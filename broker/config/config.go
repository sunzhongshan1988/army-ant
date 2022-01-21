package config

import (
	"encoding/json"
	"log"
	"os"
)

type Broker struct {
	Label           string `json:"label"`
	BrokerId        string `json:"brokerId"`
	BrokerType      string `json:"brokerType"`
	Address         string `json:"address"`
	GrpcPort        string `json:"grpcPort"`
	GraphqlPort     string `json:"graphqlPort"`
	MongodbUri      string `json:"mongodbUri"`
	MongodbDatabase string `json:"mongodbDatabase"`
}

var broker = &Broker{}

func Init() {
	broker.Label = os.Getenv("AAB_LABEL") // "system"
	broker.BrokerId = ""
	broker.BrokerType = "main"
	broker.Address = os.Getenv("AAB_ADDRESS")          // "127.0.0.1"
	broker.GrpcPort = os.Getenv("AAB_GRPC_PORT")       // "50051"
	broker.GraphqlPort = os.Getenv("AAB_GRAPHQL_PORT") // "8080"
	// mongodb uri
	// example: "mongodb://armyant:P@ssw0rd@10.11.51.152:27017/armyant_dev?authSource=admin&readPreference=primary&appname=ArmyAnt&ssl=false"
	broker.MongodbUri = os.Getenv("AAB_MONGODB_URI")
	broker.MongodbDatabase = os.Getenv("AAB_MONGODB_DATABASE") // mongodb database name

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

func GetMongodbUri() string {
	return broker.MongodbUri
}

func GetMongodbDatabase() string {
	return broker.MongodbDatabase
}
