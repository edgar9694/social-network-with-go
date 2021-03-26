package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* TweetsFollowers es la estructura con la que devolveremos los tweets*/
type TweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	Tweet          struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
	}
}
