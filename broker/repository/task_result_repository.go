package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
)

type TaskResultRepository interface {
	InsertOne(ctx context.Context, tr *model.TaskResult) (*mongo.InsertOneResult, error)
	FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.TaskResultItemsPage, error)
}

type TaskResultMongo struct {
	Database *mongo.Database
}

func (r *TaskResultMongo) InsertOne(ctx context.Context, worker *model.TaskResult) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Database.Collection("task_result").InsertOne(ctx, worker)

	if err != nil {
		log.Printf("[mgdb,save] error:%v", err)
	}

	log.Printf("[mgdb,save] info: %v", insertResult.InsertedID)

	return insertResult, err
}

func (r *TaskResultMongo) FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.TaskResultItemsPage, error) {

	result := model.TaskResultItemsPage{}

	findOptions := &options.FindOptions{}
	findOptions.SetSort(bson.M{"_id": -1})

	if page.Size > 0 {
		findOptions.SetLimit(page.Size)
		findOptions.SetSkip(page.Index * page.Size)
	}

	cur, err := r.Database.Collection("task_result").Find(context.TODO(), filter, findOptions)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,mgdb] error:%v", err)
	}

	count, err1 := r.Database.Collection("task_result").CountDocuments(ctx, filter)
	if err1 != nil {
		log.Printf("[error,mgdb] error:%v", err)
	}

	result.TotalItems = count
	result.TotalPages = int64(math.Ceil(float64(count) / float64(page.Size)))
	result.CurrentPage = page.Index

	defer cur.Close(ctx)
	if err = cur.All(ctx, &result.Items); err != nil {
		log.Printf("[mgdb, findall] error:%v", err)
	}

	log.Printf("[mgdb, findall] info: %v", "success")

	return &result, err
}
