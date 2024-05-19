package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/romanov-is-here/godeviz/internal/api"
)

func main() {
	r := mux.NewRouter()
	api.Setup(r)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./vue-app/dist/")))

	http.Handle("/", r)

	log.Println("Server is started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
