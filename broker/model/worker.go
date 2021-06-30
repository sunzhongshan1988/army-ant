package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Worker struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	BrokerId    string                 `bson:"broker_id,omitempty" json:"brokerId"`
	BrokerLink  string                 `bson:"broker_link,omitempty" json:"brokerLink"`
	WorkerId    string                 `bson:"worker_id,omitempty" json:"workerId"`
	WorkerLink  string                 `bson:"worker_link,omitempty" json:"workerLink"`
	WorkerLabel string                 `bson:"worker_label,omitempty" json:"workerLable"`
	CreateAt    *timestamppb.Timestamp `bson:"create_at" json:"createAt"`
	UpdateAt    *timestamppb.Timestamp `bson:"update_at" json:"updateAt"`
}

type WorkerItemsPage struct {
	TotalItems  int64     `json:"totalItems"`
	TotalPages  int64     `json:"totalPages"`
	CurrentPage int64     `json:"currentItems"`
	Items       []*Worker `json:"items"`
}
