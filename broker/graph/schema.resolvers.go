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

	"github.com/sunzhongshan1988/army-ant/broker/graph/generated"
	"github.com/sunzhongshan1988/army-ant/broker/graph/model"
	"github.com/sunzhongshan1988/army-ant/broker/grpc"
	pb "github.com/sunzhongshan1988/army-ant/proto/service"
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

func (r *mutationResolver) ReceiveTask(ctx context.Context, task *model.TaskInput) (*model.Task, error) {
	jsonStr, _ := json.Marshal(task)
	log.Printf("Received Task: %v", string(jsonStr))

	// Processing task DNA, mutation.
	sDec, _ := b64.StdEncoding.DecodeString(task.Dna)
	var m model.DNA
	err := json.Unmarshal([]byte(sDec), &m)
	if err != nil {
		log.Printf("[error]DNA: %v", err)
	}

	mDec, _ := b64.StdEncoding.DecodeString(task.Dna)
	var mtt model.Mutation
	err1 := json.Unmarshal([]byte(mDec), &mtt)
	if err1 != nil {
		log.Printf("[error]Mutation: %v", err)
	}

	request := &pb.TaskRequest{
		Id:   task.ID,
		Type: pb.TaskType_NOW,
		Dna: &pb.DNA{
			Cmd: &pb.Command{
				App:  m.Cmd.App,
				Args: m.Cmd.Args,
				Env:  m.Cmd.Env,
			},
			Version: m.Version,
		},
	}

	grpc.SendTask(request)

	answer := &model.Task{
		Status: 0,
		Msg:    "ok",
	}

	r.tasks = append(r.tasks, answer)
	return answer, nil
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
