package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"readit_backend/data"
	"readit_backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionsCollection *mongo.Collection
var userSubscriptionsCollection *mongo.Collection

func init() {
	client := data.GetMongoClient()
	subscriptionsCollection = client.Database("readitDB").Collection("subscriptions")
	userSubscriptionsCollection = client.Database("readitDB").Collection("user_subscriptions")
}

func CreateSubscriptionPlanHandler(w http.ResponseWriter, r *http.Request) {
	var plan models.SubscriptionPlan

	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	plan.ID = primitive.NewObjectID().Hex()

	_, err = subscriptionsCollection.InsertOne(context.TODO(), plan)
	if err != nil {
		http.Error(w, "Failed to create subscription plan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plan)
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

func PurchaseSubscription(w http.ResponseWriter, r *http.Request) {
	UserId, _ := getUserFromToken(r)
	var request struct {
		SubscriptionID string `json:"subscription_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Find the subscription plan
	var subscription models.SubscriptionPlan
	err = subscriptionsCollection.FindOne(context.TODO(), bson.M{"_id": request.SubscriptionID}).Decode(&subscription)
	if err != nil {
		http.Error(w, "Subscription not found", http.StatusNotFound)
		return
	}

	// Create user subscription
	startDate := time.Now()
	expiryDate := startDate.AddDate(0, 0, 30) // Add 30 days

	userSubscription := models.UserSubscription{
		ID:             primitive.NewObjectID().Hex(),
		UserID:         UserId.Hex(),
		SubscriptionID: request.SubscriptionID,
		StartDate:      startDate,
		ExpiryDate:     expiryDate,
	}

	_, err = userSubscriptionsCollection.InsertOne(context.TODO(), userSubscription)
	if err != nil {
		http.Error(w, "Failed to purchase subscription", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userSubscription)
}
func ViewUserSubscription(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: Failed to get user ID", http.StatusUnauthorized)
		return
	}
	log.Println("UserID:", userID)

	// Find user's active subscription
	var userSubscription models.UserSubscription
	err = userSubscriptionsCollection.FindOne(
		context.TODO(),
		bson.M{"user_id": userID, "expiry_date": bson.M{"$gte": time.Now()}},
	).Decode(&userSubscription)
	if err != nil {
		log.Println("Error retrieving subscription:", err)
		http.Error(w, "No active subscription found", http.StatusNotFound)
		return
	}

	// Calculate days left
	daysLeft := int(time.Until(userSubscription.ExpiryDate).Hours() / 24)
	if daysLeft < 0 {
		daysLeft = 0 // Handle expired subscriptions
	}

	response := struct {
		Subscription models.UserSubscription `json:"subscription"`
		DaysLeft     int                     `json:"days_left"`
		IsExpired    bool                    `json:"is_expired"`
	}{
		Subscription: userSubscription,
		DaysLeft:     daysLeft,
		IsExpired:    daysLeft <= 0,
	}

	log.Printf("Response: %+v\n", response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
