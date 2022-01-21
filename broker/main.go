package main

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/database/mgdb"
	svr "github.com/sunzhongshan1988/army-ant/broker/server"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"log"
	"sync"
)

func main() {
	log.Printf("------------Broker Started!------------")

	config.ShowVersion()

	// Initialized config
	config.Init()

	// Initialized mongo database
	mgdb.Init()
	client := mgdb.Client
	defer client.Disconnect(context.Background())

	// register broker
	brokerService := service.Broker{}
	brokerService.Register()

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
