package main

import (
	msg "./message"
	svr "./server"
	"log"
)

func main() {
	log.Printf("------------Broker Started!------------")

	msg.SendTask()

	svr.Server()
}
