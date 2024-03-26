package mongodb


type Movie struct {
	Name  string   `json:"name,omitempty" bson:"name,omitempty"`
	Amount int32   `json:"amount,omitempty" bson:"amount,omitempty"`
	Seats [][]string `json:"seats,omitempty" bson:"seats,omitempty"`
}
