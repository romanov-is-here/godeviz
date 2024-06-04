package pingcontroller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type controller struct {
}

func Setup(r *mux.Router) {
	c := controller{}

	r.HandleFunc("/api/ping", c.Ping).Methods("GET")
}

func (c *controller) Ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}
