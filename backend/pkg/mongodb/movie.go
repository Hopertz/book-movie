package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Amount int32              `json:"amount,omitempty" bson:"amount,omitempty"`
	Seats  [][]string         `json:"seats,omitempty" bson:"seats,omitempty"`
}
