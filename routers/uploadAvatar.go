package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/* UploadAvatar() sube el banner al servidor*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, _ := r.FormFile("avatar")
	var ext = strings.Split(handler.Filename, ".")[1]
	var archive = "uploads/avatars/" + IDUsuario + "." + ext

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUsuario + "." + ext
	status, err = bd.ModifyRecord(user, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
