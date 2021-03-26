package bd

import (
	"context"
	"time"

	"github.com/edgar9694/social-network-with-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(ID string, page int) ([]models.TweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	// MANERA QUE TIENE MONGODB DE UNIR DOS TABLAS
	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{"$match": bson.M{"userid": ID}})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localfield":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condition = append(condition, bson.M{"$unwind": "$tweet"})
	condition = append(condition, bson.M{"$sort": bson.M{"tweet.date": -1}})
	condition = append(condition, bson.M{"$skip": skip})
	condition = append(condition, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condition)
	var result []models.TweetsFollowers
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
