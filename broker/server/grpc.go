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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	log.Printf("[system, grpc] info: Start Grpc Server")

	lis, err := net.Listen("tcp", ":"+config.GetGrpcPort())
	if err != nil {
		log.Fatalf("[system, grpc] error: failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[system, grpc] error: failed to serve %v", err)
	}
}

// WorkerRegister implements WorkerRegister.GreeterServer
func (s *server) WorkerRegister(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	workerService := service.Worker{}
	taskService := service.Task{}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("[grpc, workerregister] info: %v", jsonStr)

	worker := &model.Worker{
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

		// Update task status
		filter1 := bson.M{"worker_id": r.WorkerId, "status": 2}
		update := bson.M{"$set": bson.M{"status": 3}}
		_, _ = taskService.UpdateOne(filter1, update)

	} else {
		worker.BrokerId = config.GetBrokerId()
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

	taskResultService := service.TaskResult{}
	taskService := service.Task{}

	res := &pb.TaskResultResponse{
		Code: 0,
		Msg:  "ok",
	}

	m := jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	jsonStr, _ := m.MarshalToString(in)
	log.Printf("[grpc, taskresult] info: %v", jsonStr)

	taskObjID, _ := primitive.ObjectIDFromHex(in.TaskId)
	tr := &model.TaskResult{
		TaskId:     taskObjID,
		InstanceID: in.InstanceId,
		BrokerId:   in.BrokerId,
		WorkerId:   in.WorkerId,
		Type:       in.Type,
		Status:     in.Status,
		Result:     in.Result,
		StartAt:    in.StartAt,
		EndAt:      in.EndAt,
	}

	// Save task result to DB
	_, err := taskResultService.InsertOne(tr)
	if err != nil {
		res.Msg = "DB error"
	}

	// Update task status
	if in.Type == 0 {
		filter1 := bson.M{"_id": taskObjID}
		update := bson.M{"$set": bson.M{"status": 2}}
		_, err2 := taskService.UpdateOne(filter1, update)
		if err2 != nil {
			res.Msg = "DB error"
		}
	}

	return res, nil
}
