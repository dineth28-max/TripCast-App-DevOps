package models

type User struct {
	ID       string `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"` // Omitted in response
	Name     string `json:"name" bson:"name"`
}
