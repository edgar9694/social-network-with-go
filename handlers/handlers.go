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
	router.HandleFunc("/modificarperfil", middlew.CheckDB(middlew.CheckJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.CheckJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/leotweet", middlew.CheckDB(middlew.CheckJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlew.CheckDB(middlew.CheckJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/subiravatar", middlew.CheckDB(middlew.CheckJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/obteneravatar", middlew.CheckDB(routers.ObtainAvatar)).Methods("GET")
	router.HandleFunc("/subirbanner", middlew.CheckDB(middlew.CheckJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtenerbanner", middlew.CheckDB(routers.ObtainBanner)).Methods("GET")

	router.HandleFunc("/altarelacion", middlew.CheckDB(middlew.CheckJWT(routers.HighRelation))).Methods("POST")
	router.HandleFunc("/bajarelacion", middlew.CheckDB(middlew.CheckJWT(routers.LowRelation))).Methods("DELETE")
	router.HandleFunc("/consultarelacion", middlew.CheckDB(middlew.CheckJWT(routers.ConsultRelation))).Methods("GET")

	router.HandleFunc("/listausuarios", middlew.CheckDB(middlew.CheckJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/leotweetsseguidores", middlew.CheckDB(middlew.CheckJWT(routers.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
