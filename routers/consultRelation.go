package routers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/* ConsultRelation() chequea si hay relacion entre 2 usuarios*/
func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUsuario
	t.UserRelationID = ID

	var resp models.ResponseConsultRelation
	status, err := bd.ConsultRelation(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
