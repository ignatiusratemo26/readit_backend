package models

type SubscriptionPlan struct {
	ID           string   `bson:"_id,omitempty" json:"id"`
	Name         string   `bson:"name" json:"name"`
	Description  string   `bson:"description" json:"description"`
	Features     []string `bson:"features" json:"features"`
	Price        float64  `bson:"price" json:"price"`
	Refreshments int      `bson:"refreshments" json:"refreshments"`
}
