package middlew

import (
	"net/http"

	"github.com/edgar9694/social-network-with-go.git/bd"
)

/* CheckDB() es el middleware que me permite saber el estado de la base de datos recibe un http handler, si la conexion a la base de datos no falla devuelve todo lo que recibe al siguiente handler*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
