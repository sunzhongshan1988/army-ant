package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	svr "github.com/sunzhongshan1988/army-ant/broker/server"
	"log"
	"sync"
	"time"
)

func main() {
	log.Printf("------------Broker Started!------------")

	// Initialized mongo database
	mongodb.Init()
	client := mongodb.GetClient()
	worker := &model.WorkerRegister{
		BrokerId:   "192.168.12.233:8088",
		BrokerLink: "192.168.12.233:8088",
		WorkerId:   "192.168.12.233:8088",
		WorkerLink: "192.168.12.233:8088",
		CreateAt:   ptypes.TimestampNow(),
		UpdateAt:   ptypes.TimestampNow(),
	}

	// Save to
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertResult, err := client.Database("armyant").Collection("worker").InsertOne(ctx, worker)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("MongoDB Save: %v", insertResult.InsertedID)
	defer client.Disconnect(context.Background())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Run graphql server
	go func() {
		svr.Graphql()
		wg.Done()
	}()

	// Run grpc server
	go func() {
		svr.Grpc()
		wg.Done()
	}()

	wg.Wait()
}
