package model

type OneKeyAnalyse struct {
	ID    int32 `bson:"_id,omitempty" json:"id"`
	Total int32 `bson:"total,omitempty" json:"total"`
}
