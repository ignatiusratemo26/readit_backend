package models

type ContactMessage struct {
	ID      string `bson:"_id,omitempty" json:"id"`
	Name    string `bson:"name" json:"name"`
	Email   string `bson:"email" json:"email"`
	Message string `bson:"message" json:"message"`
}
