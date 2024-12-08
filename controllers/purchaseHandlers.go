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

var purchaseCollection *mongo.Collection

func init() {
	client := data.GetMongoClient()
	purchaseCollection = client.Database("ndulaDB").Collection("purchases")
}

func GetPurchases(w http.ResponseWriter, r *http.Request) {
	userID, _ := getUserFromToken(r)
	var purchases []models.Purchase
	cursor, err := purchaseCollection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var purchase models.Purchase
		cursor.Decode(&purchase)
		purchases = append(purchases, purchase)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchases)
}

func PurchaseCart(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from token
	userID, err := getUserFromToken(r)
	if err != nil {
		log.Printf("Failed to extract user ID: %v", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var cartItems []models.CartItem
	cursor, err := cartCollection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {

		http.Error(w, "Failed to fetch cart items", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	// Populate cart items
	for cursor.Next(context.TODO()) {
		var cartItem models.CartItem
		if err := cursor.Decode(&cartItem); err != nil {
			log.Printf("Error decoding cart item: %v", err)
			continue
		}
		cartItems = append(cartItems, cartItem)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error while iterating cart items: %v", err)
		http.Error(w, "Failed to process cart items", http.StatusInternalServerError)
		return
	}

	// Calculate total
	var total float64
	for _, item := range cartItems {
		total += item.Product.Price * float64(item.Quantity)
	}

	// Create purchase record
	purchase := models.Purchase{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		Items:     cartItems,
		Total:     total,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	// Insert purchase into the database
	_, err = purchaseCollection.InsertOne(context.TODO(), purchase)
	if err != nil {
		log.Printf("Error inserting purchase for user %v: %v", userID.Hex(), err)
		http.Error(w, "Failed to process purchase", http.StatusInternalServerError)
		return
	}

	// Clear the cart
	_, err = cartCollection.DeleteMany(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		log.Printf("Error clearing cart for user %v: %v", userID.Hex(), err)
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}

	// Respond with the purchase details
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(purchase); err != nil {
		log.Printf("Error encoding response for user %v: %v", userID.Hex(), err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
