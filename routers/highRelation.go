package routers

import (
	"net/http"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/* HighRelation() realiza el registro de la relacion entre usuarios */
func HighRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUsuario
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
