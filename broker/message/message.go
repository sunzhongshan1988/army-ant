package message

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "../../proto/service"
)

const (
	workerAddress = "localhost:50052"
)

func SendTask() {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(workerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	request := &pb.TaskRequest{
		Id:   uuid.New().String(),
		Type: pb.TaskType_NOW,
		Cmd: &pb.Command{
			App:  "adb",
			Args: []string{"version"},
			Env:  []string{""},
		},
	}
	defer cancel()
	r, err2 := c.SendTask(ctx, request)
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("Greeting: %s", jsonStr)
}
