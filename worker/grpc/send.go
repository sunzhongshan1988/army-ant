package grpc

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/sunzhongshan1988/army-ant/worker/config"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

const (
	brokerAddress = "localhost:50051"
)

func Register() {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.RegisterRequest{
		Auth:       "#shdk687dHHhiJHDHDHH",
		WorkerType: pb.WorkerType_IDC,
		WorkerLink: config.GetAddress() + ":" + strconv.Itoa(int(config.GetPort())),
		Content:    "",
		CreateAt:   ptypes.TimestampNow(),
	}

	r, err2 := c.WorkerRegister(ctx, request)
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("Broker Response: %s", jsonStr)

	config.SetBrokerId(r.BrokerId)
	config.SetWorkerId(r.WorkerId)
	config.SetBrokerLink(r.BrokerLink)

}

func TaskResult(result string, status int32, start *timestamppb.Timestamp, end *timestamppb.Timestamp) {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.TaskResultRequest{
		WorkerId: config.GetWorkerId(),
		BrokerId: config.GetBrokerId(),
		Status:   status,
		Result:   result,
		StartAt:  start,
		EndAt:    end,
	}

	r, err2 := c.TaskResult(ctx, request)
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("Broker Response: %s", jsonStr)

}
