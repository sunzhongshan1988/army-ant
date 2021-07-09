package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BrokerId string             `bson:"broker_id,omitempty" json:"brokerId"`
	WorkerId string             `bson:"worker_id,omitempty" json:"workerId"`
	Type     int                `bson:"type" json:"type"`
	Cron     string             `bson:"cron,omitempty" json:"cron"`
	DNA      DNA                `bson:"dna" json:"dna"`
	Mutation Mutation           `bson:"mutation" json:"mutation"`
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
