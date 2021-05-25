package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkerRegister struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty"`
	BrokerId   string                 `bson:"broker_id,omitempty"`
	BrokerLink string                 `bson:"broker_link,omitempty"`
	WorkerId   string                 `bson:"worker_id,omitempty"`
	WorkerLink string                 `bson:"worker_link,omitempty"`
	CreateAt   *timestamppb.Timestamp `bson:"create_at"`
	UpdateAt   *timestamppb.Timestamp `bson:"update_at"`
}
