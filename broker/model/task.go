package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Task struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	InstanceId string                 `bson:"instance_id,omitempty" json:"instanceId"`
	BrokerId   string                 `bson:"broker_id,omitempty" json:"brokerId"`
	WorkerId   string                 `bson:"worker_id,omitempty" json:"workerId"`
	EntryId    int32                  `bson:"entry_id,omitempty" json:"entryId"`
	Type       int64                  `bson:"type" json:"type"`
	Status     int32                  `bson:"status" json:"status"`
	Cron       string                 `bson:"cron,omitempty" json:"cron"`
	DNA        string                 `bson:"dna" json:"dna"`
	Mutation   string                 `bson:"mutation" json:"mutation"`
	CreateAt   *timestamppb.Timestamp `bson:"create_at" json:"createAt"`
	UpdateAt   *timestamppb.Timestamp `bson:"update_at" json:"updateAt"`
}

type TaskItemsPage struct {
	TotalItems  int64   `json:"totalItems"`
	TotalPages  int64   `json:"totalPages"`
	CurrentPage int64   `json:"currentItems"`
	Items       []*Task `json:"items"`
}

type Command struct {
	App  string   `json:"app"`
	Args []string `json:"args"`
	Env  []string `json:"env"`
}
type DNA struct {
	Cmd     *Command `json:"cmd"`
	Version string   `json:"version"`
}

type CommandMutation struct {
	Args []string `json:"args"`
	Env  []string `json:"env"`
}
type Mutation struct {
	Cmd     *CommandMutation `json:"cmd"`
	Version string           `json:"version"`
}
