package grpc

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

const (
	workerAddress = "localhost:50052"
)

func SendTask(request *pb.TaskRequest, workerId string) {
	workerService := service.Worker{}

	filter := bson.M{"worker_id": workerId}
	worker, err := workerService.FindOne(filter)
	if err != nil {
		log.Printf("[grpc, sendtask] error: %v", err)
	}
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(worker.WorkerLink, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Printf("[grpc, sendtask] error: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err2 := c.SendTask(ctx, request)
	if err2 != nil {
		log.Printf("[grpc, sendtask] error: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("[grpc, sendtask] info: %v", jsonStr)
}
