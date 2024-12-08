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

var subscriptionsCollection *mongo.Collection

func init() {
	client := data.GetMongoClient()
	subscriptionsCollection = client.Database("ndulaDB").Collection("subscriptions")
}

func GetSubscriptionPlansHandler(w http.ResponseWriter, r *http.Request) {
	cursor, err := subscriptionsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve plans", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var plans []models.SubscriptionPlan
	if err = cursor.All(context.TODO(), &plans); err != nil {
		http.Error(w, "Failed to parse plans", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(plans)
}
