package server

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Status:      1,
		CreateAt:    in.CreateAt,
		UpdateAt:    timestamppb.Now(),
	}

	// Query Database
	filter := bson.M{"worker_link": in.WorkerLink, "worker_label": in.WorkerLabel}
	r, _ := workerService.FindOne(filter)
	if r != nil {
		worker.BrokerId = r.BrokerId
		worker.WorkerId = r.WorkerId

		// Update task status, if task is running(code 1) then set status as suspend(code: 3)
		filter1 := bson.M{"worker_id": r.WorkerId, "status": 1}
		update1 := bson.M{"$set": bson.M{"status": 3}}
		_, _ = taskService.UpdateMany(filter1, update1)

		// Update worker status and update time
		filter2 := bson.M{"worker_id": r.WorkerId}
		update2 := bson.M{"$set": bson.M{"status": 1, "update_at": worker.UpdateAt}}
		_, _ = workerService.UpdateOne(filter2, update2)

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
		TaskName:   in.TaskName,
		TaskRemark: in.TaskRemark,
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
	if in.Type == 0 || in.Type == 2 {
		filter1 := bson.M{"_id": taskObjID}
		update := bson.M{"$set": bson.M{"status": 2}}
		_, err2 := taskService.UpdateOne(filter1, update)
		if err2 != nil {
			res.Msg = "DB error"
		}
	}

	return res, nil
}
