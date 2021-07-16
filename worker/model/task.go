package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Input struct {
	TaskID     string
	InstanceID string
	App        string
	Args       []string
	Env        []string
}

type CommandResult struct {
	TaskID     string
	InstanceID string
	Out        string
	Status     int32
	StartAt    *timestamp.Timestamp
	EndAt      *timestamp.Timestamp
}
