package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/sunzhongshan1988/army-ant/broker/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	characters  []*model.Character
	tasks       []*model.Task
	brokerItems []*model.BrokerItems
}
