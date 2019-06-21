package api

import (
	"net/http"

	"{{.ModName}}"
	log "github.com/sirupsen/logrus"
)

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	log.Info("Request get all users")
	u, err := userInfoSrv.GetAll()
	if err != nil {
		log.WithError(err).Error("request get all user failed")
		// TODO: your must define what kind of error to indicate http code is 400 or 500
		respondWithJSON(w, 500, err) 
		return
	}

	res := struct {
		Data []{{.ProjName}}.User `json:"data"`
	} {
		Data: u,
	}

	respondWithJSON(w, 200, res)
}
