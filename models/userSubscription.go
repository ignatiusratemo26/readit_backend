package models

import "time"

type UserSubscription struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	UserID         string    `bson:"user_id" json:"user_id"`
	SubscriptionID string    `bson:"subscription_id" json:"subscription_id"`
	StartDate      time.Time `bson:"start_date" json:"start_date"`
	ExpiryDate     time.Time `bson:"expiry_date" json:"expiry_date"`
}
