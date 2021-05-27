package grpc

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	pb "github.com/sunzhongshan1988/army-ant/proto/service"
	pf "github.com/sunzhongshan1988/army-ant/worker/performer"
	"log"
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
