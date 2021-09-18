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
	Total    int64     `json:"total"`
	PageSize int64     `json:"pageSize"`
	Current  int64     `json:"current"`
	Items    []*Worker `json:"items"`
}
