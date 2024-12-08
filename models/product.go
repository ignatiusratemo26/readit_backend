package models

type Product struct {
	ID          string  `bson:"_id,omitempty" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
	Price       float64 `bson:"price" json:"price"`
	Category    string  `bson:"category" json:"category"`
}
