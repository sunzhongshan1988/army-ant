package config

type Worker struct {
	workerId   string
	workerType string
	brokerId   string
	brokerLink string
	address    string
	port       int32
}

var worker = &Worker{}

func Init() {
	worker.workerId = ""
	worker.workerType = "g"
	worker.brokerId = ""
	worker.brokerLink = "localhost:50051"
	worker.address = "127.0.0.1"
	worker.port = 50052
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

func GetAddress() string {
	return worker.address
}

func GetPort() int32 {
	return worker.port
}
