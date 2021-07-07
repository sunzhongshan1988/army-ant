package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskResult struct {
	ID       primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	BrokerId string                 `bson:"broker_id,omitempty" json:"brokerId"`
	WorkerId string                 `bson:"worker_id,omitempty" json:"workerId"`
	Status   int32                  `bson:"status,omitempty" json:"status"`
	Result   string                 `bson:"result,omitempty" json:"result"`
	StartAt  *timestamppb.Timestamp `bson:"start_at" json:"startAt"`
	EndAt    *timestamppb.Timestamp `bson:"end_at" json:"endAt"`
}
type TaskResultItemsPage struct {
	TotalItems  int64         `json:"totalItems"`
	TotalPages  int64         `json:"totalPages"`
	CurrentPage int64         `json:"currentItems"`
	Items       []*TaskResult `json:"items"`
}