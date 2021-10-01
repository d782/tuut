package routers

import (
	"encoding/json"
	"net/http"

	"github.com/d782/tuut/bd"
	"github.com/d782/tuut/models"
)

/* sign in user */

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password is too short", 400)
	}

	_, found, _ := bd.UserExist(t.Email)
	if found == true {
		http.Error(w, "Email already exist!", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrió un error"+err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "failed to insert information", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
