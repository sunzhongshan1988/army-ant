package grpc

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	pb "github.com/sunzhongshan1988/army-ant/proto/service"
	"log"
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
