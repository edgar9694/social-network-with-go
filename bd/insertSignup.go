package bd

import (
	"context"
	"time"

	"github.com/edgar9694/social-network-with-go.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertSignup() es la parada final con la base de datos para insertar los datos del usuario */
func InsertSignup(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
