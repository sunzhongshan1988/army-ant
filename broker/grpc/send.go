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

func SendTask(request *pb.TaskRequest, workerId string) (entryId int32, status int32) {
	workerService := service.Worker{}

	filter := bson.M{"worker_id": workerId}
	worker, err := workerService.FindOne(filter)
	if err != nil {
		log.Printf("[grpc, sendtask] error: %v", err)
		return 0, 1
	}
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(worker.WorkerLink, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Printf("[grpc, sendtask] error: %v", err1)
		return 0, 1
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err2 := c.Task(ctx, request)
	if err2 != nil {
		log.Printf("[grpc, sendtask] error: %v", err2)
		return 0, 1
	}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("[grpc, sendtask] info: %v", jsonStr)
	return r.EntryId, 0
}

func StopTask(request *pb.StopTaskRequest) (*pb.StopTaskResponse, error) {
	workerService := service.Worker{}

	res := &pb.StopTaskResponse{
		Status: 0,
		Msg:    "ok",
	}

	filter := bson.M{"worker_id": request.WorkerId}
	worker, err := workerService.FindOne(filter)
	if err != nil {
		log.Printf("[grpc, stoptask] error: %v", err)
		res.Msg = "db error"
		return res, err
	}

	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(worker.WorkerLink, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Printf("[grpc, stoptask] error: %v", err1)
		res.Msg = "connect worker error"
		return res, err1
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err2 := c.StopTask(ctx, request)
	if err2 != nil {
		log.Printf("[grpc, stoptask] error: %v", err2)
		res.Msg = "send task error"
		return res, err2
	}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("[grpc, stoptask] info: %v", jsonStr)

	return r, nil

}
