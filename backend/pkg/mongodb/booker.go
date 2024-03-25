package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SelectedSeat struct {
	Row  int `json:"row,omitempty" bson:"row,omitempty"`
	Seat int `json:"seat,omitempty" bson:"seat,omitempty"`
}

type Booker struct {
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Seat      []SelectedSeat     `json:"seat,omitempty" bson:"seat,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

