package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/* Email valor de Email usado en todos los EndPoints*/
var Email string

/* IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/* ProcessToken() proceso el token para extraer sus valores */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("SolucionesFinancieras")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, finded, _ := bd.CheckUserExists(claims.Email)
		if finded {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, finded, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
