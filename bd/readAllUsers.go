package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/edgar9694/social-network-with-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ReadAllUsers() lee todo los usuarios registrados en el sistema, si se escribe "R" en quienes */
func ReadAllUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var result []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"firstName": bson.M{"$regex": `(?i)` + search}, // ESTE REGEX ES PARA BUSCAR SIN REVISAR SI ES MAYUSCULA O MINUSCULA
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	var finded, included bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		included = false
		finded, _ = ConsultRelation(r)
		if userType == "new" && !finded {
			included = true
		}

		if userType == "follow" && finded {
			included = true
		}

		if r.UserRelationID == ID {
			included = false
		}

		if included {
			s.Password = ""
			s.Biography = ""
			s.Address = ""
			s.WebSite = ""
			s.Banner = ""
			s.Email = ""

			result = append(result, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	cur.Close(ctx)
	return result, true
}
