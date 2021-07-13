package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/sunzhongshan1988/army-ant/broker/config"
	"github.com/sunzhongshan1988/army-ant/broker/graph/generated"
	"github.com/sunzhongshan1988/army-ant/broker/grpc"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"github.com/sunzhongshan1988/army-ant/broker/service"
	pb "github.com/sunzhongshan1988/army-ant/proto/service"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) Add(ctx context.Context, character model.CharacterInput) (*model.Character, error) {
	charac := &model.Character{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Name:  character.Name,
		Likes: character.Likes,
	}

	r.characters = append(r.characters, charac)
	return charac, nil
}

func (r *mutationResolver) ReceiveTask(ctx context.Context, task *model.TaskInput) (*model.TaskResponse, error) {
	jsonStr, _ := json.Marshal(task)
	log.Printf("Received Task: %v", string(jsonStr))

	// Processing task DNA, mutation.
	sDec, _ := b64.StdEncoding.DecodeString(task.Dna)
	var m model.DNA
	err := json.Unmarshal([]byte(sDec), &m)
	if err != nil {
		log.Printf("[error]DNA: %v", err)
	}

	mDec, _ := b64.StdEncoding.DecodeString(task.Mutation)
	var mtt model.Mutation
	err1 := json.Unmarshal([]byte(mDec), &mtt)
	if err1 != nil {
		log.Printf("[error]Mutation: %v", err)
	}

	request := &pb.TaskRequest{
		Id:       task.InstanceID,
		Type:     task.Type,
		Cron:     task.Cron,
		BrokerId: config.GetBrokerId(),
		WorkerId: task.WorkerID,
		Dna: &pb.DNA{
			Cmd: &pb.Command{
				App:  m.Cmd.App,
				Args: mtt.Cmd.Args,
				Env:  mtt.Cmd.Env,
			},
			Version: mtt.Version,
		},
	}

	entryId, sendStatus := grpc.SendTask(request, task.WorkerID)

	res := &model.TaskResponse{
		Status: 0,
		Msg:    "ok",
	}
	if sendStatus == 1 {
		res.Status = 1
		res.Msg = "send task to worker error"
	}

	if sendStatus == 0 {
		taskDb := &model.Task{
			InstanceId: task.InstanceID,
			BrokerId:   config.GetBrokerId(),
			WorkerId:   task.WorkerID,
			EntryId:    entryId,
			Type:       task.Type,
			Status:     sendStatus,
			Cron:       task.Cron,
			DNA:        task.Dna,
			Mutation:   task.Mutation,
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

func (r *mutationResolver) StopTask(ctx context.Context, task *model.StopTaskInput) (*model.StopTaskResponse, error) {
	jsonStr, _ := json.Marshal(task)
	log.Printf("[graphql, stoptask] info: %v", string(jsonStr))

	res := &model.StopTaskResponse{
		Status: 1,
		Msg:    "error",
	}

	req := &pb.StopTaskRequest{
		Id:       task.InstanceID,
		BrokerId: task.BrokerID,
		WorkerId: task.WorkerID,
		EntryId:  0,
	}
	taskService := service.Task{}
	filter := bson.M{"instance_id": task.InstanceID, "worker_id": task.WorkerID}
	dbtask, err := taskService.FindOne(filter)
	if err != nil {
		res.Msg = "query db error"
		return res, err
	}
	req.EntryId = dbtask.EntryId
	grpcres, err1 := grpc.StopTask(req)
	if err1 != nil {
		res.Msg = "send to worker error"
		return res, err1
	}

	filter1 := bson.M{"worker_id": task.WorkerID}
	update := bson.M{"$set": bson.M{"status": 1}}
	_, err2 := taskService.UpdateOne(filter1, update)
	if err2 != nil {
		res.Msg = "update db error"
		return res, err
	}

	res.Status = 1
	res.Msg = grpcres.Msg
	return res, nil
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return r.characters, nil
}

func (r *queryResolver) Search(ctx context.Context, name string) (*model.Character, error) {
	charName := strings.ToLower(name)
	for _, x := range r.characters {
		if strings.Contains(strings.ToLower(x.Name), charName) {
			return x, nil
		}
	}
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
		TotalPages:  dbRes.TotalPages,
		TotalItems:  dbRes.TotalItems,
		CurrentPage: dbRes.CurrentPage,
		Items:       dbRes.Items,
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
		TotalPages:  dbRes.TotalPages,
		TotalItems:  dbRes.TotalItems,
		CurrentPage: dbRes.CurrentPage,
		Items:       dbRes.Items,
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
		TotalPages:  dbRes.TotalPages,
		TotalItems:  dbRes.TotalItems,
		CurrentPage: dbRes.CurrentPage,
		Items:       dbRes.Items,
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
