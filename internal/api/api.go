package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/romanov-is-here/godeviz/internal/api/controllers/graphcontroller"
	"github.com/romanov-is-here/godeviz/internal/api/controllers/pingcontroller"
)

func Setup(r *mux.Router, path string) {
	r.HandleFunc("/api/hello", getHello).Methods("GET")

	graphcontroller.Setup(r, path)
	pingcontroller.Setup(r)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello world"))
	if err != nil {
		return
	}
}
