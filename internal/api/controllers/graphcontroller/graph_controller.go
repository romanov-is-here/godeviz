package graphcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/romanov-is-here/godeviz/internal/api/controllers/graphcontroller/dto"
	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
	"github.com/romanov-is-here/godeviz/internal/graph/lister"
	"github.com/romanov-is-here/godeviz/internal/graph/painter"
)

type graphController struct {
}

func SetupGraphController(r *mux.Router) {
	c := graphController{}

	r.HandleFunc("/api/graph", c.GetGraph).Methods("GET")
}

func (c *graphController) GetGraph(w http.ResponseWriter, r *http.Request) {
	path, ok := fromQuery(w, r, "path", true)
	if !ok {
		return
	}

	err, g := lister.GetGraph(path)
	if err != nil {
		http.Error(w, "Failed to make a graph", http.StatusBadRequest)
		return
	}

	countEdge := 0
	countNode := 0
	nodeIds := make(map[string]string)
	refd := painter.GetRefdFromHome(g)

	newNodeId := func() string {
		countNode++
		return "node" + strconv.Itoa(countNode)
	}

	newEdgeId := func() string {
		countEdge++
		return "edge" + strconv.Itoa(countEdge)
	}

	outGraph := dto.Graph{
		Nodes: make(map[string]*dto.Node),
		Edges: make(map[string]*dto.Edge),
	}

	// TODO questionable code below
	// Collect nodes
	for _, pck := range g.Packs {
		if _, ok := refd[pck.Id()]; !ok {
			continue
		}
		node := dto.Node{
			Name:   pck.Name(),
			IsHome: pck.IsHome,
			Color:  getColor(pck),
		}
		id := newNodeId()
		nodeIds[pck.Id()] = id
		outGraph.Nodes[id] = &node
	}

	// Collect hits
	for _, pck := range g.Packs {
		if _, ok := refd[pck.Id()]; !ok || !pck.IsHome {
			continue
		}
		nodeId, ok := nodeIds[pck.Id()]
		if !ok {
			continue
		}
		node, ok := outGraph.Nodes[nodeId]
		if !ok {
			continue
		}
		for _, imp := range pck.Imports() {
			if _, ok := refd[imp.Id()]; !ok {
				continue
			}
			node.OutDeps++
			depId, ok := nodeIds[imp.Id()]
			if !ok {
				continue
			}
			depNode, ok := outGraph.Nodes[depId]
			if !ok {
				continue
			}
			depNode.InDeps++
		}
		//node.InDeps = pck.Count()
		//node.OutDeps = pck.ImportsCount()
	}

	// Collect edges
	for _, srcPack := range g.Packs {
		if !srcPack.IsHome {
			continue
		}
		srcNodeId, ok := nodeIds[srcPack.Id()]
		if !ok {
			continue
		}
		for _, imp := range srcPack.Imports() {
			destPackId, ok := nodeIds[imp.Id()]
			if !ok {
				continue
			}
			destPack, ok := g.Packs[imp.Id()]
			if !ok {
				continue
			}
			edge := dto.Edge{
				Source: srcNodeId,
				Target: destPackId,
				Color:  getColor(destPack),
			}
			edgeId := newEdgeId()
			outGraph.Edges[edgeId] = &edge
		}
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

func getColor(p *depgraph.PackageInfo) string {
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
