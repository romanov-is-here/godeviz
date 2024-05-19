package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", getHello).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./vue-app/dist/")))

	http.Handle("/", r)

	log.Println("Сервер запущен на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getHello(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("Hello world"))
	if err != nil {
		return
	}
	//json.NewEncoder(w).Encode(todos)
}
