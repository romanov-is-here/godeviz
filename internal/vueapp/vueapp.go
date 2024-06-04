package vueapp

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	vapp "github.com/romanov-is-here/godeviz/vue-app"
)

func Setup(r *mux.Router) {
	stripped, err := fs.Sub(vapp.Content, "dist")
	if err != nil {
		log.Fatalln(err)
	}
	r.PathPrefix("/").Handler(http.FileServer(http.FS(stripped)))
	http.Handle("/", r)
}
