package config

import (
	"encoding/json"
	"log"
	"os"
)

type Worker struct {
	WorkerId   string `json:"workerId"`
	WorkerType string `json:"workerType"`
	BrokerId   string `json:"brokerId"`
	BrokerLink string `json:"brokerLink"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	Label      string `json:"label"`
	Version    string `json:"version"`
}

var worker = &Worker{}

func Init() {
	worker.WorkerId = ""
	worker.WorkerType = "g"
	worker.BrokerId = ""
	worker.BrokerLink = os.Getenv("AAW_BROKER_LINK")
	worker.Address = os.Getenv("AAW_ADDRESS")
	worker.Port = os.Getenv("AAW_PORT")
	worker.Label = os.Getenv("AAW_LABEL")
	worker.Version = "0.0.1"

	jsonStr, _ := json.Marshal(worker)
	log.Printf("[config, init] info: %v", string(jsonStr))
}

func GetWorkerId() string {
	return worker.WorkerId
}
func SetWorkerId(id string) {
	worker.WorkerId = id
}

func GetWorkerType() string {
	return worker.WorkerType
}

func GetBrokerId() string {
	return worker.BrokerId
}
func SetBrokerId(id string) {
	worker.BrokerId = id
}

func GetBrokerLink() string {
	return worker.BrokerLink
}
func SetBrokerLink(link string) {
	worker.BrokerLink = link
}

func GetAddress() string {
	return worker.Address
}

func GetPort() string {
	return worker.Port
}

func GetWorkerLink() string {
	return worker.Address + ":" + worker.Port
}

func GetLabel() string {
	return worker.Label
}

func GetVersion() string {
	return worker.Version
}
