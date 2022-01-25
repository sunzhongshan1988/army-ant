package main

import (
	"github.com/sunzhongshan1988/army-ant/worker/config"
	"github.com/sunzhongshan1988/army-ant/worker/cronmod"
	"github.com/sunzhongshan1988/army-ant/worker/grpc"
	svr "github.com/sunzhongshan1988/army-ant/worker/server"
	"log"
)

func main() {
	log.Printf("------------Worker Started!------------")
	// print version information.
	config.ShowVersion()

	config.Init()

	cronmod.Init()

	grpc.Register()

	svr.Grpc()

}
