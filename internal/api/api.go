package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/romanov-is-here/godeviz/internal/api/controllers/graphcontroller"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/api/hello", getHello).Methods("GET")

	graphcontroller.SetupGraphController(r)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello world"))
	if err != nil {
		return
	}
}
