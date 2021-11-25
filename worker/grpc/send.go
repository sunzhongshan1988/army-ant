package grpc

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sunzhongshan1988/army-ant/worker/config"
	"github.com/sunzhongshan1988/army-ant/worker/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"

	pb "github.com/sunzhongshan1988/army-ant/proto/service"
)

func Register() {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(config.GetBrokerLink(), grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Printf("[grpc, register] error: did not connect: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.RegisterRequest{
		Auth:          "#shdk687dHHhiJHDHDHH",
		WorkerType:    pb.WorkerType_IDC,
		WorkerLink:    config.GetWorkerLink(),
		Content:       "",
		WorkerLabel:   config.GetLabel(),
		WorkerVersion: config.GetVersion(),
		CreateAt:      timestamppb.Now(),
	}

	r, err2 := c.WorkerRegister(ctx, request)
	if err2 != nil {
		log.Printf("[grpc, register] error: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("[grpc, register] info: %s", jsonStr)

	config.SetBrokerId(r.BrokerId)
	config.SetWorkerId(r.WorkerId)
	config.SetBrokerLink(r.BrokerLink)

}

func TaskResult(commandResult *model.CommandResult) {
	// Set up a connection to the broker.
	conn, err1 := grpc.Dial(config.GetBrokerLink(), grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil {
		log.Printf("[grpc, taskresult] error: %v", err1)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the broker and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.TaskResultRequest{
		WorkerId:   config.GetWorkerId(),
		BrokerId:   config.GetBrokerId(),
		TaskId:     commandResult.TaskID,
		TaskName:   commandResult.TaskName,
		TaskRemark: commandResult.TaskRemark,
		InstanceId: commandResult.InstanceID,
		Status:     commandResult.Status,
		Type:       commandResult.Type,
		Result:     commandResult.Out,
		StartAt:    commandResult.StartAt,
		EndAt:      commandResult.EndAt,
	}

	r, err2 := c.TaskResult(ctx, request)
	if err2 != nil {
		log.Printf("[grpc, taskresult] error: %v", err2)
	}
	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(r)
	log.Printf("[grpc, taskresult] info: %s", jsonStr)

}
