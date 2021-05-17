package message

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

const (
	workerAddress = "localhost:50052"
)

func SendTask(request *pb.TaskRequest) {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(workerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	r, err2 := c.SendTask(ctx, request)
	if err2 != nil {
		log.Fatalf("Error: could not greet: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("Worker Response: %s", jsonStr)
}
