package repository

import (
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

type TaskRepository interface {
	AnalyseTaskStatus(ctx context.Context, pipeline mongo.Pipeline) (*[]model.AnalyseTaskStatus, error)
	FindOne(ctx context.Context, filter bson.M) (*model.Task, error)
	FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.TaskItemsPage, error)
	InsertOne(ctx context.Context, tr *model.Task) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error)
}

type TaskMongo struct {
	Database *mongo.Database
}

func (r *TaskMongo) AnalyseTaskStatus(ctx context.Context, pipeline mongo.Pipeline) (*[]model.AnalyseTaskStatus, error) {
	var result []model.AnalyseTaskStatus

	cur, err := r.Database.Collection("task").Aggregate(ctx, pipeline)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[mgdb, aggregate] error: %v", err)
	}

	defer cur.Close(ctx)
	if err = cur.All(ctx, &result); err != nil {
		log.Printf("[mgdb, aggregate] error:%v", err)
	}

	log.Printf("[mgdb, aggregate] info: %v", "success")

	return &result, err
}

func (r *TaskMongo) FindOne(ctx context.Context, filter bson.M) (*model.Task, error) {

	var result model.Task

	err := r.Database.Collection("task").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[mgdb, findone] error: %v", err)
	}

	log.Printf("[mgdb, findone] info: %v", "success")

	return &result, err
}

func (r *TaskMongo) FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.TaskItemsPage, error) {

	result := model.TaskItemsPage{}

	findOptions := &options.FindOptions{}
	findOptions.SetSort(bson.M{"_id": -1})

	if page.Size > 0 {
		findOptions.SetLimit(page.Size)
		findOptions.SetSkip(page.Index * page.Size)
	}

	cur, err := r.Database.Collection("task").Find(context.TODO(), filter, findOptions)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,mgdb] error:%v", err)
	}

	count, err1 := r.Database.Collection("task").CountDocuments(ctx, filter)
	if err1 != nil {
		log.Printf("[error,mgdb] error:%v", err)
	}

	result.Total = count
	result.PageSize = page.Size
	result.Current = page.Index

	defer cur.Close(ctx)
	if err = cur.All(ctx, &result.Items); err != nil {
		log.Printf("[mgdb, findall] error:%v", err)
	}

	log.Printf("[mgdb, findall] info: %v", "success")

	return &result, err
}

func (r *TaskMongo) InsertOne(ctx context.Context, worker *model.Task) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Database.Collection("task").InsertOne(ctx, worker)

	if err != nil {
		log.Printf("[mgdb,save] error:%v", err)
	}

	log.Printf("[mgdb,save] info: %v", insertResult.InsertedID)

	return insertResult, err
}

func (r *TaskMongo) UpdateOne(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error) {

	updateResult, err := r.Database.Collection("task").UpdateOne(ctx, filter, data)

	if err != nil {
		log.Printf("[mgdb,updateone] error:%v", err)
	}

	log.Printf("[mgdb,updateone] info: %v", updateResult.ModifiedCount)

	return updateResult, err
}

func (r *TaskMongo) UpdateMany(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error) {

	updateResult, err := r.Database.Collection("task").UpdateMany(ctx, filter, data)

	if err != nil {
		log.Printf("[mgdb,updateone] error:%v", err)
	}

	log.Printf("[mgdb,updateone] info: %v", updateResult.ModifiedCount)

	return updateResult, err
}
