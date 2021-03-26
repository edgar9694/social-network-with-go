package bd

import (
	"context"
	"time"

	"github.com/edgar9694/social-network-with-go.git/models"
)

/* InsertRelation() graba la relación en la base de datos */
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
