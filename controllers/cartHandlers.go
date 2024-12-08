package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"readit_backend/data"
	"readit_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection

func init() {
	client := data.GetMongoClient()
	cartCollection = client.Database("readitDB").Collection("cart")
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, _ := getUserFromToken(r)
	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cartItem.ID = primitive.NewObjectID()
	cartItem.UserID = userID
	_, err := cartCollection.InsertOne(context.TODO(), cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cartItem)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID, _ := getUserFromToken(r)
	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := cartCollection.DeleteOne(context.TODO(), bson.M{"_id": cartItem.ID, "user_id": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cartItem)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	userID, _ := getUserFromToken(r)
	var cartItems []models.CartItem

	cursor, err := cartCollection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var cartItem models.CartItem
		cursor.Decode(&cartItem)
		cartItems = append(cartItems, cartItem)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cartItems)
}
