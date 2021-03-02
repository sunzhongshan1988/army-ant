package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "../../proto/service"
	pf "../performer"
)

const (
	serverPort = ":50052"
)

// server is used to implement server.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
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
		App:  in.Cmd.App,
		Args: in.Cmd.Args,
		Env:  in.Cmd.Env,
	}
	output := pf.Standard(input)
	log.Printf("cmd-out: %v", output.StdoutPipeOut)
	log.Printf("cmd-err: %v", output.StdoutPipeErr)

	res := &pb.TaskResponse{
		Status: 1,
		Msg:    "ok",
	}
	return res, nil
}

func Server() {
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
