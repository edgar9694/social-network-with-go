package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/* UploadBanner() sube el banner al servidor */
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, _ := r.FormFile("banner")
	var ext = strings.Split(handler.Filename, ".")[1]
	var archive = "uploads/banners/" + IDUsuario + "." + ext

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir el banner! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar el banner! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUsuario + "." + ext
	status, err = bd.ModifyRecord(user, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
