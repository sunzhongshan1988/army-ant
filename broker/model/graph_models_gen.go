// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BrokerPageResponse struct {
	TotalItems  int64     `json:"totalItems"`
	TotalPages  int64     `json:"totalPages"`
	CurrentPage int64     `json:"currentPage"`
	Items       []*Broker `json:"items"`
}

type Character struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Likes int64  `json:"likes"`
}

type CharacterInput struct {
	Name  string `json:"name"`
	Likes int64  `json:"likes"`
}

type GetBrokerItemsInput struct {
	Index int64 `json:"index"`
	Size  int64 `json:"size"`
}

type GetWorkerItemsInput struct {
	Index int64 `json:"index"`
	Size  int64 `json:"size"`
}

type Task struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
}

type TaskInput struct {
	ID       string `json:"id"`
	BrokerID string `json:"broker_id"`
	WorkerID string `json:"worker_id"`
	Type     string `json:"type"`
	Dna      string `json:"dna"`
	Mutation string `json:"mutation"`
}

type WorkerPageResponse struct {
	TotalItems  int64     `json:"totalItems"`
	TotalPages  int64     `json:"totalPages"`
	CurrentPage int64     `json:"currentPage"`
	Items       []*Worker `json:"items"`
}