package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/robfig/cron/v3"
	"github.com/sunzhongshan1988/army-ant/worker/config"
	"github.com/sunzhongshan1988/army-ant/worker/cronmod"
	"github.com/sunzhongshan1988/army-ant/worker/model"
	pf "github.com/sunzhongshan1988/army-ant/worker/performer"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

// server is used to implement server.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func Grpc() {
	// Start server
	log.Printf("[system, grpc] info: Start gRPC Server")
	lis, err := net.Listen("tcp", ":"+config.GetPort())
	if err != nil {
		log.Fatalf("[system, grpc] error: failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[system, grpc] error: failed to serve %v", err)
	}
}

// Task SendTask implements SendTask.GreeterServer
func (s *server) Task(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("[grpc, task] info: %v", jsonStr)

	input := &model.Input{
		TaskID:     in.TaskId,
		TaskName:   in.TaskName,
		TaskRemark: in.TaskRemark,
		InstanceID: in.InstanceId,
		Type:       in.Type,
		App:        in.Dna.Cmd.App,
		Args:       in.Dna.Cmd.Args,
		Env:        in.Dna.Cmd.Env,
		Dir:        in.Dna.Cmd.Dir,
	}

	res := &pb.TaskResponse{
		Status:  0,
		EntryId: 0,
		Msg:     "ok",
	}

	switch in.Type {
	case 0:
		go pf.Standard(*input)
	case 1:
		entryId, err := cronmod.AddFunc(in.Cron, func() { pf.Standard(*input) })
		if err != nil {
			res.Status = 1
			res.Msg = err.Error()
		}
		res.Status = 0
		res.EntryId = int32(entryId)
		res.Msg = "ok"
	case 2:
		go pf.Standard(*input)
	case 3:
		go pf.Standard(*input)
	}

	return res, nil
}

// StopTask implements StopTask.GreeterServer
func (s *server) StopTask(ctx context.Context, in *pb.StopTaskRequest) (*pb.StopTaskResponse, error) {
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("[grpc, stoptask] info: %v", jsonStr)

	res := &pb.StopTaskResponse{
		Status: 1,
		Msg:    "ok",
	}

	cronmod.Remove(cron.EntryID(in.EntryId))

	return res, nil
}

// KillTask implements KillTask.GreeterServer
func (s *server) KillTask(ctx context.Context, in *pb.KillTaskRequest) (*pb.KillTaskResponse, error) {
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("[grpc, killtask] info: %v", jsonStr)

	res := &pb.KillTaskResponse{
		Status: 1,
		Msg:    "ok",
	}

	if ok := pf.Kill(in.TaskId); !ok {
		res.Status = 2
		res.Msg = "task not running"
	}

	return res, nil
}
