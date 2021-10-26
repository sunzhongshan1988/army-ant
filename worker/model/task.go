package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Input struct {
	TaskID     string
	TaskName   string
	TaskRemark string
	InstanceID string
	Type       int64
	App        string
	Args       []string
	Env        []string
	Dir        string
}

type CommandResult struct {
	TaskID     string
	TaskName   string
	TaskRemark string
	InstanceID string
	Out        string
	Type       int64
	Status     int32
	StartAt    *timestamp.Timestamp
	EndAt      *timestamp.Timestamp
}
