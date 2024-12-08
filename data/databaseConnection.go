package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongo() {
	var err error

	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo connection error: ", err)
	}

}
func GetMongoClient() *mongo.Client {
	InitMongo()
	if mongoClient == nil {
		log.Fatal("MongoDB client not initialized. Call InitMongo() first.")
	}
	return mongoClient
}
