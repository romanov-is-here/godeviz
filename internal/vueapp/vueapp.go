package vueapp

import (
	"net/http"

	"github.com/gorilla/mux"

	vapp "github.com/romanov-is-here/godeviz/vue-app"
)

func Setup(r *mux.Router) {
	r.PathPrefix("/").Handler(http.FileServer(http.FS(vapp.Content)))
	http.Handle("/", r)
}
