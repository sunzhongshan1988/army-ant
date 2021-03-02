package message

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "../../proto/service"
)

const (
	brokerAddress = "localhost:50051"
)

func Message() {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	request := &pb.RegisterRequest{
		Auth:       "#shdk687dHHhiJHDHDHH",
		WorkerType: pb.WorkerType_IDC,
		Content:    "",
		CreateAt:   ptypes.TimestampNow(),
	}
	defer cancel()
	r, err2 := c.WorkerRegister(ctx, request)
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
