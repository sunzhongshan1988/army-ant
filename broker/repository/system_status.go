package repository

import (
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type SystemStatusRepository interface {
	FindOne(ctx context.Context, filter bson.M) (*model.Broker, error)
}

type SystemStatusMongo struct {
	Client *mongo.Client
}
