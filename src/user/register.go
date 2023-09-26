package user

import (
	"encoding/json"
	"golang_ws_template/src/clients"
	"golang_ws_template/src/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clients.MongoDBUserCreate(creds)

}
