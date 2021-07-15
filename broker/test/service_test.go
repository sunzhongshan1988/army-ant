package test

import (
	"context"
	"encoding/json"
	"github.com/sunzhongshan1988/army-ant/broker/database/mgdb"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestBrokerService(t *testing.T) {
	// Initialized mongo database
	mongodb.Init()
	client := mongodb.Client
	defer client.Disconnect(context.Background())

	var brokerService = service.BrokerService{}
	page := &model.PageableRequest{
		Index: 0,
		Size:  10,
	}
	res, _ := brokerService.FindAll(bson.M{}, page)
	jsonStr, _ := json.Marshal(res)
	log.Printf("Result: %v", string(jsonStr))
}
