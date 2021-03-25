package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/edgar9694/social-network-with-go.git/middlew"
	"github.com/edgar9694/social-network-with-go.git/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers() seteo mi puerto, el Handler y pongo a escuchar al servidor*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckDB(middlew.CheckJWT(routers.Profile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
