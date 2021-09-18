package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/graph/generated"
	"github.com/sunzhongshan1988/army-ant/broker/grpc"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	pb "github.com/sunzhongshan1988/army-ant/proto/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mutationResolver) ReceiveTask(ctx context.Context, task *model.TaskInput) (*model.StdResponse, error) {
	jsonStr, _ := json.Marshal(task)
	log.Printf("[graphql, receivedtask] info: %v", string(jsonStr))

	// Processing task DNA, mutation.
	sDec, _ := b64.StdEncoding.DecodeString(task.Dna)
	var m model.DNA
	err := json.Unmarshal([]byte(sDec), &m)
	if err != nil {
		log.Printf("[graphql, receivedtask] error: DNA %v", err)
	}

	mDec, _ := b64.StdEncoding.DecodeString(task.Mutation)
	var mtt model.Mutation
	err1 := json.Unmarshal([]byte(mDec), &mtt)
	if err1 != nil {
		log.Printf("[graphql, receivedtask] error: Mutation %v", err)
	}

	taskID := primitive.NewObjectID()
	request := &pb.TaskRequest{
		TaskName:   task.Name,
		InstanceId: task.InstanceID,
		TaskId:     taskID.Hex(),
		Type:       task.Type,
		Cron:       task.Cron,
		BrokerId:   config.GetBrokerId(),
		WorkerId:   task.WorkerID,
		TaskRemark: task.Remark,
		Dna: &pb.DNA{
			Cmd: &pb.Command{
				App:  m.Cmd.App,
				Args: mtt.Cmd.Args,
				Env:  mtt.Cmd.Env,
				Dir:  mtt.Cmd.Dir,
			},
			Version: mtt.Version,
		},
	}

	entryId, err3 := grpc.SendTask(request, task.WorkerID)

	res := &model.StdResponse{
		Status: 0,
		Msg:    taskID.Hex(),
	}
	if err3 != nil {
		res.Status = 1
		res.Msg = "send task to worker error"
	}

	if err3 == nil {
		taskDb := &model.Task{
			ID:         taskID,
			Name:       task.Name,
			InstanceId: task.InstanceID,
			BrokerId:   config.GetBrokerId(),
			WorkerId:   task.WorkerID,
			EntryId:    entryId,
			Type:       task.Type,
			Status:     1,
			Cron:       task.Cron,
			DNA:        task.Dna,
			Mutation:   task.Mutation,
			Remark:     task.Remark,
			CreateAt:   ptypes.TimestampNow(),
			UpdateAt:   ptypes.TimestampNow(),
		}

		taskService := service.Task{}
		_, err2 := taskService.InsertOne(taskDb)
		if err2 != nil {
			res.Status = 1
			res.Msg = "Broker DB error!"

			return res, nil
		}
	}

	//r.tasks = append(r.tasks, res)
	return res, nil
}

func (r *mutationResolver) StopTask(ctx context.Context, task *model.TaskInstanceInput) (*model.StdResponse, error) {
	jsonStr, _ := json.Marshal(task)
	log.Printf("[graphql, stoptask] info: %v", string(jsonStr))

	res := &model.StdResponse{
		Status: 1,
		Msg:    "error",
	}

	req := &pb.StopTaskRequest{
		Id:       task.TaskID,
		BrokerId: task.BrokerID,
		WorkerId: "",
		EntryId:  0,
	}

	taskObjID, _ := primitive.ObjectIDFromHex(task.TaskID)
	filter := bson.M{"_id": taskObjID, "status": 1}
	taskService := service.Task{}
	dbtask, err := taskService.FindOne(filter)
	if err != nil {
		res.Msg = "query db error"
		return res, err
	}
	req.WorkerId = dbtask.WorkerId
	req.EntryId = dbtask.EntryId
	grpcres, err1 := grpc.StopTask(req)
	if err1 != nil {
		res.Msg = "send to worker error"
		return res, err1
	}

	filter1 := bson.M{"_id": dbtask.ID}
	update := bson.M{"$set": bson.M{"status": 2}}
	_, err2 := taskService.UpdateOne(filter1, update)
	if err2 != nil {
		res.Msg = "update db error"
		return res, err
	}

	res.Status = 0
	res.Msg = grpcres.Msg
	return res, nil
}

func (r *mutationResolver) RetryTask(ctx context.Context, task *model.TaskInstanceInput) (*model.StdResponse, error) {
	jsonStr, _ := json.Marshal(task)
	log.Printf("[graphql, retrytask] info: %v", string(jsonStr))

	res := &model.StdResponse{
		Status: 1,
		Msg:    "error",
	}

	taskObjID, _ := primitive.ObjectIDFromHex(task.TaskID)
	filter := bson.M{"_id": taskObjID}
	taskService := service.Task{}
	dbTask, err2 := taskService.FindOne(filter)
	if err2 != nil {
		res.Msg = "query db error"
		return res, err2
	}

	if dbTask.Status == 1 {
		res.Msg = "wait task finish"
		return res, err2
	}

	// Processing task DNA, mutation.
	sDec, _ := b64.StdEncoding.DecodeString(dbTask.DNA)
	var m model.DNA
	err := json.Unmarshal([]byte(sDec), &m)
	if err != nil {
		log.Printf("[graphql, retrytask] error: DNA %v", err)
	}

	mDec, _ := b64.StdEncoding.DecodeString(dbTask.Mutation)
	var mtt model.Mutation
	err1 := json.Unmarshal([]byte(mDec), &mtt)
	if err1 != nil {
		log.Printf("[graphql, retrytask] error: Mutation %v", err)
	}

	req := &pb.TaskRequest{
		InstanceId: dbTask.InstanceId,
		TaskId:     dbTask.ID.Hex(),
		Type:       dbTask.Type,
		Cron:       dbTask.Cron,
		BrokerId:   dbTask.BrokerId,
		WorkerId:   dbTask.WorkerId,
		Dna: &pb.DNA{
			Cmd: &pb.Command{
				App:  m.Cmd.App,
				Args: mtt.Cmd.Args,
				Env:  mtt.Cmd.Env,
				Dir:  mtt.Cmd.Dir,
			},
			Version: mtt.Version,
		},
	}

	entryId, err3 := grpc.SendTask(req, dbTask.WorkerId)
	if err3 != nil {
		res.Msg = "update db error"
		return res, err3
	}

	filter1 := bson.M{"_id": dbTask.ID}
	update := bson.M{"$set": bson.M{"status": 1, "entry_id": entryId}}
	_, err2 = taskService.UpdateOne(filter1, update)
	if err2 != nil {
		res.Msg = "update db error"
		return res, err2
	}

	res.Status = 0
	res.Msg = "ok"
	return res, nil
}

func (r *queryResolver) GetSystemStatus(ctx context.Context) (*model.SystemStatusResponse, error) {
	taskService := service.Task{}

	pipeline := mongo.Pipeline{
		{{"$group", bson.D{{"_id", "$status"}, {"total", bson.D{{"$sum", 1}}}}}},
	}
	_, _ = taskService.AnalyseTaskStatus(pipeline)
	return nil, nil
}

func (r *queryResolver) GetBrokerItems(ctx context.Context, page *model.GetBrokerItemsInput) (*model.BrokerPageResponse, error) {
	brokerService := service.Broker{}

	pg := &model.PageableRequest{
		Index: page.Index,
		Size:  page.Size,
	}
	dbRes, _ := brokerService.FindAll(bson.M{}, pg)

	res := &model.BrokerPageResponse{
		Total:    dbRes.Total,
		PageSize: dbRes.PageSize,
		Current:  dbRes.Current,
		Items:    dbRes.Items,
	}
	return res, nil
}

func (r *queryResolver) GetWorkerItems(ctx context.Context, page *model.GetWorkerItemsInput) (*model.WorkerPageResponse, error) {
	workerService := service.Worker{}

	pg := &model.PageableRequest{
		Index: page.Index,
		Size:  page.Size,
	}
	dbRes, _ := workerService.FindAll(bson.M{}, pg)

	res := &model.WorkerPageResponse{
		Total:    dbRes.Total,
		PageSize: dbRes.PageSize,
		Current:  dbRes.Current,
		Items:    dbRes.Items,
	}
	return res, nil
}

func (r *queryResolver) GetTaskItems(ctx context.Context, page *model.GetTaskItemsInput) (*model.TaskPageResponse, error) {
	taskService := service.Task{}

	pg := &model.PageableRequest{
		Index: page.Index,
		Size:  page.Size,
	}
	dbRes, _ := taskService.FindAll(bson.M{}, pg)

	res := &model.TaskPageResponse{
		Total:    dbRes.Total,
		PageSize: dbRes.PageSize,
		Current:  dbRes.Current,
		Items:    dbRes.Items,
	}
	return res, nil
}

func (r *queryResolver) GetTaskResultItems(ctx context.Context, page *model.GetTaskResultItemsInput) (*model.TaskResultPageResponse, error) {
	taskResultService := service.TaskResult{}

	pg := &model.PageableRequest{
		Index: page.Index,
		Size:  page.Size,
	}
	dbRes, _ := taskResultService.FindAll(bson.M{}, pg)

	res := &model.TaskResultPageResponse{
		Total:    dbRes.Total,
		PageSize: dbRes.PageSize,
		Current:  dbRes.Current,
		Items:    dbRes.Items,
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
