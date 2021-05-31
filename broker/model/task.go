package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskResult struct {
	ID       primitive.ObjectID     `bson:"_id,omitempty"`
	BrokerId string                 `bson:"broker_id,omitempty"`
	WorkerId string                 `bson:"worker_id,omitempty"`
	Status   int32                  `bson:"status,omitempty"`
	Result   string                 `bson:"result,omitempty"`
	StartAt  *timestamppb.Timestamp `bson:"start_at"`
	EndAt    *timestamppb.Timestamp `bson:"end_at"`
}
