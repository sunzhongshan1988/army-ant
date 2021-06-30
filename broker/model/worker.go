package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Worker struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty" json:""`
	BrokerId    string                 `bson:"broker_id,omitempty" json:""`
	BrokerLink  string                 `bson:"broker_link,omitempty" json:""`
	WorkerId    string                 `bson:"worker_id,omitempty" json:""`
	WorkerLink  string                 `bson:"worker_link,omitempty" json:""`
	WorkerLabel string                 `bson:"worker_label,omitempty" json:""`
	CreateAt    *timestamppb.Timestamp `bson:"create_at" json:""`
	UpdateAt    *timestamppb.Timestamp `bson:"update_at" json:""`
}

type WorkerItemsPage struct {
	TotalItems  int64     `json:"totalItems"`
	TotalPages  int64     `json:"totalPages"`
	CurrentPage int64     `json:"currentItems"`
	Items       []*Worker `json:"items"`
}
