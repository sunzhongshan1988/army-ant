package model

import "go.mongodb.org/mongo-driver/bson"

type PageableRequest struct {
	Index int64
	Size  int64
}

type PageableResponse struct {
	TotalItems  int64    `json:"totalItems"`
	TotalPages  int64    `json:"totalPages"`
	CurrentPage int64    `json:"currentItems"`
	Items       []bson.M `json:"items"`
}
