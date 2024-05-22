package graphcontroller

import (
	"github.com/romanov-is-here/godeviz/internal/api/controllers/graphcontroller/dto"
	"github.com/romanov-is-here/godeviz/internal/graph/lister"
)

type graphService struct{}

func (s *graphService) getGraph(path string) (*dto.Graph, error) {

	err, g := lister.GetGraph(path)
	if err != nil {
		return nil, err
	}

	outGraph := dto.Graph{
		Nodes: make(map[string]*dto.Node),
		Edges: make(map[string]*dto.Edge),
	}

	// Collect nodes
	for _, pck := range g.Packs {
		node := dto.Node{
			Name:    pck.Name,
			IsHome:  pck.IsHome,
			Color:   getColor(pck),
			InDeps:  pck.FanIn,
			OutDeps: pck.FanOut,
		}
		outGraph.Nodes[pck.Id] = &node
	}

	// Collect edges
	for _, srcPack := range g.Packs {
		for _, imp := range srcPack.Imports {
			destPack, ok := g.Packs[imp.Id]
			if !ok {
				continue
			}
			edge := dto.Edge{
				Source: srcPack.Id,
				Target: imp.Id,
				Color:  getColor(destPack),
			}
			edgeId := srcPack.Id + "->" + imp.Id
			outGraph.Edges[edgeId] = &edge
		}
	}

	return &outGraph, nil
}
