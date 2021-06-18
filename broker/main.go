package main

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	svr "github.com/sunzhongshan1988/army-ant/broker/server"
	"log"
	"sync"
)

func main() {
	log.Printf("------------Broker Started!------------")

	// Initialized mongo database
	mongodb.Init()
	client := mongodb.Client
	defer client.Disconnect(context.Background())

	// Initialized config
	config.Init()

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
