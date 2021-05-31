package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	pf "github.com/sunzhongshan1988/army-ant/worker/performer"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

const (
	serverPort = ":50052"
)

// server is used to implement server.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func Grpc() {
	// Start server
	log.Printf("--Start Server")
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

// SendTask implements SendTask.GreeterServer
func (s *server) SendTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("Task: %v", jsonStr)

	input := pf.Input{
		App:  in.Dna.Cmd.App,
		Args: in.Dna.Cmd.Args,
		Env:  in.Dna.Cmd.Env,
	}

	res := &pb.TaskResponse{
		Status: 1,
		Msg:    "ok",
	}

	go pf.Standard(input)

	return res, nil
}
