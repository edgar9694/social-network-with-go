package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/models"
)

/*  InsertTweet() permite grabar el tweet en la base de datos*/
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Formato Incorrecto "+err.Error(), 400)
		return
	}
	record := models.SaveTweet{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(record)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
