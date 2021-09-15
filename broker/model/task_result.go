package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TaskResult Status 	0 - success, 1 - error
type TaskResult struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	TaskId     primitive.ObjectID     `bson:"task_id,omitempty" json:"taskId"`
	InstanceID string                 `bson:"instance_id,omitempty" json:"instanceId"`
	BrokerId   string                 `bson:"broker_id,omitempty" json:"brokerId"`
	WorkerId   string                 `bson:"worker_id,omitempty" json:"workerId"`
	Status     int32                  `bson:"status,omitempty" json:"status"`
	Type       int64                  `bson:"type,omitempty" json:"type"`
	Result     string                 `bson:"result,omitempty" json:"result"`
	StartAt    *timestamppb.Timestamp `bson:"start_at" json:"startAt"`
	EndAt      *timestamppb.Timestamp `bson:"end_at" json:"endAt"`
}

type TaskResultItemsPage struct {
	TotalItems  int64         `json:"totalItems"`
	TotalPages  int64         `json:"totalPages"`
	CurrentPage int64         `json:"currentItems"`
	Items       []*TaskResult `json:"items"`
}
