package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Broker Status: 0 - stateless, 1 - online, 2 - offline
type Broker struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	BrokerId    string                 `bson:"broker_id,omitempty" json:"brokerId"`
	BrokerLink  string                 `bson:"broker_link,omitempty" json:"brokerLink"`
	BrokerLabel string                 `bson:"broker_label,omitempty" json:"brokerLabel"`
	Status      int64                  `bson:"status,omitempty" json:"status"`
	CreateAt    *timestamppb.Timestamp `bson:"create_at" json:"createAt"`
	UpdateAt    *timestamppb.Timestamp `bson:"update_at" json:"updateAt"`
}

type BrokerItemsPage struct {
	Total    int64     `json:"total"`
	PageSize int64     `json:"pageSize"`
	Current  int64     `json:"current"`
	Items    []*Broker `json:"items"`
}
