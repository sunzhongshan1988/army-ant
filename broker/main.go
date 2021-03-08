package main

import (
	svr "github.com/sunzhongshan1988/army-ant/broker/server"
	"log"
	"sync"
)

func main() {
	log.Printf("------------Broker Started!------------")

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		svr.Graphql()
		wg.Done()
	}()

	go func() {
		svr.Grpc()
		wg.Done()
	}()

	wg.Wait()
}
