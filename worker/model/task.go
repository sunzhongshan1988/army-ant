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
	Output     string
	Error      string
	Type       int64
	Status     int32
	Duration   int64
	StartAt    *timestamp.Timestamp
	EndAt      *timestamp.Timestamp
}
