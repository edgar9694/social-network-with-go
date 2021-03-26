package bd

import (
	"context"
	"log"
	"time"

	"github.com/edgar9694/social-network-with-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ReadTweets lee los tweets de un perfil*/
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var result []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var record models.ReturnTweets
		err := cursor.Decode(&record)
		if err != nil {
			return result, false
		}

		result = append(result, &record)
	}
	return result, true
}
