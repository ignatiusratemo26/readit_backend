package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"readit_backend/data"
	"readit_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productsCollection *mongo.Collection

func init() {
	client := data.GetMongoClient()
	productsCollection = client.Database("ndulaDB").Collection("products")
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	cursor, err := productsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		http.Error(w, "Failed to parse products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
