package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TaskResult Status 	0 - success, 1 - error
type TaskResult struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	TaskId     primitive.ObjectID     `bson:"task_id,omitempty" json:"taskId"`
	TaskName   string                 `bson:"task_name,omitempty" json:"taskName"`
	TaskRemark string                 `bson:"task_remark,omitempty" json:"taskRemark"`
	InstanceID string                 `bson:"instance_id,omitempty" json:"instanceId"`
	BrokerId   string                 `bson:"broker_id,omitempty" json:"brokerId"`
	WorkerId   string                 `bson:"worker_id,omitempty" json:"workerId"`
	Status     int32                  `bson:"status,omitempty" json:"status"`
	Type       int64                  `bson:"type,omitempty" json:"type"`
	Error      string                 `bson:"error,omitempty" json:"error"`
	Output     string                 `bson:"output,omitempty" json:"output"`
	Duration   int64                  `bson:"duration,omitempty" json:"duration"`
	StartAt    *timestamppb.Timestamp `bson:"start_at" json:"startAt"`
	EndAt      *timestamppb.Timestamp `bson:"end_at" json:"endAt"`
}

type TaskResultItemsPage struct {
	Total    int64         `json:"total"`
	PageSize int64         `json:"pageSize"`
	Current  int64         `json:"current"`
	Items    []*TaskResult `json:"items"`
}
