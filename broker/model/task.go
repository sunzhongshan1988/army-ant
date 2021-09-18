package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Task Status 	0 - null, 1 - running, 2 - finish, 3 - suspend
// Task Type 	0 - manual, 1 - crontab
type Task struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Name       string                 `bson:"name,omitempty" json:"name"`
	InstanceId string                 `bson:"instance_id,omitempty" json:"instanceId"`
	BrokerId   string                 `bson:"broker_id,omitempty" json:"brokerId"`
	WorkerId   string                 `bson:"worker_id,omitempty" json:"workerId"`
	EntryId    int32                  `bson:"entry_id,omitempty" json:"entryId"`
	Type       int64                  `bson:"type" json:"type"`
	Status     int32                  `bson:"status" json:"status"`
	Cron       string                 `bson:"cron,omitempty" json:"cron"`
	DNA        string                 `bson:"dna" json:"dna"`
	Mutation   string                 `bson:"mutation" json:"mutation"`
	Remark     string                 `bson:"remark" json:"remark"`
	CreateAt   *timestamppb.Timestamp `bson:"create_at" json:"createAt"`
	UpdateAt   *timestamppb.Timestamp `bson:"update_at" json:"updateAt"`
}

type TaskItemsPage struct {
	Total    int64   `json:"total"`
	PageSize int64   `json:"pageSize"`
	Current  int64   `json:"current"`
	Items    []*Task `json:"items"`
}

type Command struct {
	App  string   `json:"app"`
	Args []string `json:"args"`
	Env  []string `json:"env"`
	Dir  string   `json:"dir"`
}
type DNA struct {
	Cmd     *Command `json:"cmd"`
	Version string   `json:"version"`
}

type CommandMutation struct {
	Args []string `json:"args"`
	Env  []string `json:"env"`
	Dir  string   `json:"dir"`
}
type Mutation struct {
	Cmd     *CommandMutation `json:"cmd"`
	Version string           `json:"version"`
}

type AnalyseTaskStatus struct {
	ID    int32 `bson:"_id,omitempty" json:"id"`
	Total int32 `bson:"total,omitempty" json:"total"`
}
