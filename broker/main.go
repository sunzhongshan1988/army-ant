package main

import (
	"context"
	mongo "github.com/sunzhongshan1988/army-ant/broker/database/mongodb"
	svr "github.com/sunzhongshan1988/army-ant/broker/server"
	"log"
	"sync"
)

func main() {
	log.Printf("------------Broker Started!------------")

	// Initialized mongo database
	client := mongo.Init()
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
