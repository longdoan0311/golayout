package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"{{.ModName}}"
)

// Cfg configuration struct of api package
type Cfg struct {
	Port string
	UserInfoSrv {{.ProjName}}.UserInfoService
}

var port string
var userInfoSrv {{.ProjName}}.UserInfoService

// Init initilize api package, pass config into
func Init(c Cfg) {
	port = c.Port
	userInfoSrv = c.UserInfoSrv
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("middleware %s", r.URL)
		h.ServeHTTP(w, r)
	})
}

// Run start http server
func Run() error {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/", handleHome).Methods("GET")

	user := r.PathPrefix("/users").Subrouter()
	user.HandleFunc("", handleGetUser).Methods("GET")

	http.Handle("/", r)
	log.Infof("{{.ProjName}} service is running at port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func respondWithJSON(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.WithError(err).WithField("data", data).Error("failed to marshal data")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(resp)
	return
}
