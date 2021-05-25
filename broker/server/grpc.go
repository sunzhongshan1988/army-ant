package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

const (
	serverPort = ":50051"
)

// server is used to implement server.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// WorkerRegister implements WorkerRegister.GreeterServer
func (s *server) WorkerRegister(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var workerService = service.WorkerService{}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("Worker Register: %v", jsonStr)

	brokerId := uuid.New().String()
	workerId := uuid.New().String()

	worker := &model.WorkerRegister{
		BrokerId:   brokerId,
		BrokerLink: "192.168.12.233:8088",
		WorkerId:   workerId,
		WorkerLink: in.WorkerLink,
		CreateAt:   in.CreateAt,
		UpdateAt:   ptypes.TimestampNow(),
	}

	// Save to
	_, _ = workerService.InsertOne(worker)

	res := &pb.RegisterResponse{
		BrokerId:   worker.BrokerId,
		WorkerId:   worker.WorkerId,
		BrokerLink: worker.BrokerLink,
		CreateAt:   worker.CreateAt,
	}
	return res, nil
}

func Grpc() {
	// Start server
	log.Printf("--Start Grpc Server")

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
