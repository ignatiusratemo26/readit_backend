package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Purchase struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Items     []CartItem         `bson:"items" json:"items"`
	Total     float64            `bson:"total" json:"total"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}
