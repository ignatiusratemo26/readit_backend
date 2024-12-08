package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	Product  Product            `bson:"product" json:"product"`
	Quantity int                `bson:"quantity" json:"quantity"`
}
