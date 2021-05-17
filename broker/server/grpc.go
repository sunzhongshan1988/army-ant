package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
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
	brokerId := uuid.New().String()
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("Worker Register: %v", jsonStr)

	res := &pb.RegisterResponse{
		BrokerId:   brokerId,
		WorkerId:   uuid.New().String(),
		BrokerLink: "192.168.12.233:8088",
		CreateAt:   ptypes.TimestampNow(),
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
