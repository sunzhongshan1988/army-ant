package model

type PageableRequest struct {
	Index int64
	Size  int64
}

type BrokerPageResponse struct {
	TotalItems  int64            `json:"totalItems"`
	TotalPages  int64            `json:"totalPages"`
	CurrentPage int64            `json:"currentItems"`
	Items       []BrokerRegister `json:"items"`
}
