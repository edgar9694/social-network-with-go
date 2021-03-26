package bd

import (
	"context"
	"time"

	"github.com/edgar9694/social-network-with-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModifyRecord() modifica un usuario en la BD*/
func ModifyRecord(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	recordList := make(map[string]interface{})
	if len(u.FirstName) > 0 {
		recordList["firstName"] = u.FirstName
	}
	if len(u.LastName) > 0 {
		recordList["lastName"] = u.LastName
	}
	recordList["birthDate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		recordList["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		recordList["banner"] = u.Banner
	}
	if len(u.Address) > 0 {
		recordList["address"] = u.Address
	}
	if len(u.Biography) > 0 {
		recordList["biography"] = u.Biography
	}
	if len(u.WebSite) > 0 {
		recordList["webSite"] = u.WebSite
	}

	updtString := bson.M{
		"$set": recordList,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
