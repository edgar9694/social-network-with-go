package middlew

import (
	"net/http"

	"github.com/edgar9694/social-network-with-go.git/routers"
)

func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
