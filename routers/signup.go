package routers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/*SignUp() es la funcion para crear en la BD el registro de usuario*/
func SignUp(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t) // el BODY de un json es de tipo STREAM es decir que solo se puede usar una vez
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	_, exists, _ := bd.CheckUserExists(t.Email)
	if exists {
		http.Error(w, "Ya existe un usuario registrado con este correo", 400)
		return
	}

	_, status, err := bd.InsertSignup(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
