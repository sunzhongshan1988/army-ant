package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"go.mongodb.org/mongo-driver/bson"
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
	log.Printf("--Start Grpc Server")

	lis, err := net.Listen("tcp", ":"+config.GetGrpcPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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

	worker := &model.WorkerRegister{
		BrokerId:    "",
		BrokerLink:  config.GetGrpcLink(),
		WorkerId:    "",
		WorkerLink:  in.WorkerLink,
		WorkerLabel: in.WorkerLabel,
		CreateAt:    in.CreateAt,
		UpdateAt:    ptypes.TimestampNow(),
	}

	// Query Database
	filter := bson.M{"worker_link": in.WorkerLink, "worker_label": in.WorkerLabel}
	r, _ := workerService.FindOne(filter)
	if r != nil {
		worker.BrokerId = r.BrokerId
		worker.WorkerId = r.WorkerId
	} else {
		worker.BrokerId = uuid.New().String()
		worker.WorkerId = uuid.New().String()

		// Save worker's information to DB
		_, _ = workerService.InsertOne(worker)
	}

	res := &pb.RegisterResponse{
		BrokerId:   worker.BrokerId,
		WorkerId:   worker.WorkerId,
		BrokerLink: worker.BrokerLink,
		CreateAt:   worker.CreateAt,
	}
	return res, nil
}

// TaskResult implements WorkerRegister.GreeterServer
func (s *server) TaskResult(ctx context.Context, in *pb.TaskResultRequest) (*pb.TaskResultResponse, error) {

	var taskResultService = service.TaskResultService{}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("Task Result: %v", jsonStr)

	tr := &model.TaskResult{
		BrokerId: in.BrokerId,
		WorkerId: in.WorkerId,
		Status:   in.Status,
		Result:   in.Result,
		StartAt:  in.StartAt,
		EndAt:    in.EndAt,
	}

	// Save to
	_, _ = taskResultService.InsertOne(tr)

	res := &pb.TaskResultResponse{
		Code: 0,
		Msg:  "ok",
	}
	return res, nil
}
