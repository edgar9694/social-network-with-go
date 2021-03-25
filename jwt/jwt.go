package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edgar9694/social-network-with-go.git/models"
)

func GenerateJWT(t models.User) (string, error) {
	secretKey := []byte("SolucionesFinancieras")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"firstname": t.FirstName,
		"lastname":  t.LastName,
		"birthday":  t.BirthDate,
		"biography": t.Biography,
		"address":   t.Address,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
