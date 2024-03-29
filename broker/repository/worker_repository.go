package repository

import (
	"context"
	"github.com/sunzhongshan1988/army-ant/broker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type WorkerRepository interface {
	InsertOne(ctx context.Context, worker *model.Worker) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error)
	FindOne(ctx context.Context, filter bson.M) (*model.Worker, error)
	FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.WorkerItemsPage, error)
}

type WorkerMongo struct {
	Database *mongo.Database
}

func (r *WorkerMongo) InsertOne(ctx context.Context, worker *model.Worker) (*mongo.InsertOneResult, error) {

	insertResult, err := r.Database.Collection("worker").InsertOne(ctx, worker)
	if err != nil {
		log.Printf("[mgdb,error]: %v", err)
	}

	log.Printf("[mgdb,save]: %v", insertResult.InsertedID)

	return insertResult, err
}

func (r *WorkerMongo) UpdateOne(ctx context.Context, filter bson.M, data bson.M) (*mongo.UpdateResult, error) {

	updateResult, err := r.Database.Collection("worker").UpdateOne(ctx, filter, data)

	if err != nil {
		log.Printf("[mgdb,updateone] error:%v", err)
	}

	log.Printf("[mgdb,updateone] info: %v", updateResult.ModifiedCount)

	return updateResult, err
}

func (r *WorkerMongo) FindOne(ctx context.Context, filter bson.M) (*model.Worker, error) {

	var result model.Worker

	err := r.Database.Collection("worker").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,db]%v", err)
	}

	log.Printf("MongoDB FindOne: %v", "success")

	return &result, err
}

func (r *WorkerMongo) FindAll(ctx context.Context, filter bson.M, page *model.PageableRequest) (*model.WorkerItemsPage, error) {

	result := model.WorkerItemsPage{}

	findOptions := &options.FindOptions{}
	findOptions.SetSort(bson.M{"_id": -1})
	if page.Size > 0 {
		findOptions.SetLimit(page.Size)
		findOptions.SetSkip(page.Index * page.Size)
	}

	cur, err := r.Database.Collection("worker").Find(context.TODO(), filter, findOptions)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		log.Printf("[error,db]%v", err)
	}

	count, err1 := r.Database.Collection("worker").CountDocuments(ctx, filter)
	if err1 != nil {
		log.Printf("[error,db]%v", err)
	}

	result.Total = count
	result.PageSize = page.Size
	result.Current = page.Index

	defer cur.Close(ctx)
	if err = cur.All(ctx, &result.Items); err != nil {
		log.Printf("[error,db]%v", err)
	}

	log.Printf("[mongo,query]: %v", "success")

	return &result, err
}
