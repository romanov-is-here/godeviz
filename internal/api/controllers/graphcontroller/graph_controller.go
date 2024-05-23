package graphcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
)

type graphController struct {
	service *graphService
}

func SetupGraphController(r *mux.Router) {
	c := graphController{service: &graphService{}}

	r.HandleFunc("/api/graph", c.GetGraph).Methods("GET")
}

func (c *graphController) GetGraph(w http.ResponseWriter, r *http.Request) {
	path, ok := strFromQuery(w, r, "path", true)
	if !ok {
		return
	}

	filter, ok := getFilter(w, r)
	if !ok {
		return
	}

	outGraph, err := c.service.getGraph(path, filter)

	if err != nil {
		http.Error(w, "Failed to make a graph", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(outGraph)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Serialization Error", http.StatusInternalServerError)
		return
	}
}

func getColor(p *depgraph.Package) string {
	if p.IsHome {
		return "brown"
	}
	if p.IsPlatform {
		return "purple"
	}
	if p.IsStandard {
		return "#0bd63e"
	}
	if p.IsOuter {
		return "lightskyblue"
	}
	return "lightgrey"
}
