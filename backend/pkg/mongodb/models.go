package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Models struct {
	Users *mongo.Collection
	Movie  *mongo.Collection
}

func NewModels(db *mongo.Client) *Models {
	return &Models{
		Users: GetCollection(db, "booker"),
		Movie: GetCollection(db, "movie"),
	}
}

func GetCollection(db *mongo.Client, collectionName string) *mongo.Collection {
	collection := db.Database("appointment").Collection(collectionName)
	return collection
}
