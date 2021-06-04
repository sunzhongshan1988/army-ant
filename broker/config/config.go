package config

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
