package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"readit_backend/data"
	"readit_backend/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var contactsCollection *mongo.Collection

func init() {
	client := data.GetMongoClient()
	contactsCollection = client.Database("readitDB").Collection("contacts")
}

func SubmitContactHandler(w http.ResponseWriter, r *http.Request) {
	var message models.ContactMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = contactsCollection.InsertOne(context.TODO(), message)
	if err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact message submitted successfully"})
}
