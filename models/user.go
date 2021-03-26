package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User es el modelo de usuario en la base de MondoDB */
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Address   string             `bson:"address" json:"address,omitempty"`
	WebSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
