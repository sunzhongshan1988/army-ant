package config

type Worker struct {
	workerId   string
	workerType string
	brokerId   string
	brokerLink string
	address    string
	port       string
	label      string
}

var worker = &Worker{}

func Init() {
	worker.workerId = ""
	worker.workerType = "g"
	worker.brokerId = ""
	worker.brokerLink = "localhost:50051"
	worker.address = "127.0.0.1"
	worker.port = "50052"
	worker.label = "worker01"
}

func GetWorkerId() string {
	return worker.workerId
}
func SetWorkerId(id string) {
	worker.workerId = id
}

func GetWorkerType() string {
	return worker.workerType
}

func GetBrokerId() string {
	return worker.brokerId
}
func SetBrokerId(id string) {
	worker.brokerId = id
}

func GetBrokerLink() string {
	return worker.brokerLink
}
func SetBrokerLink(link string) {
	worker.brokerLink = link
}

func GetAddress() string {
	return worker.address
}

func GetPort() string {
	return worker.port
}

func GetWorkerLink() string {
	return worker.address + ":" + worker.port
}

func GetLabel() string {
	return worker.label
}
