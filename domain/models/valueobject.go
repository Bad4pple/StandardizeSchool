package models

type Option struct {
	OptionID CodeID  `json:"option_id" bson:"option_id"`
	Price    float64 `json:"price" bson:"price"`
}
