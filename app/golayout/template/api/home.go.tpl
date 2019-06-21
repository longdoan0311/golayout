package api

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{{.ProjName}} service is running")
}
