package graphcontroller

import (
	"github.com/romanov-is-here/godeviz/internal/api/controllers/graphcontroller/dto"
	"github.com/romanov-is-here/godeviz/internal/graph/lister"
	"github.com/romanov-is-here/godeviz/internal/graph/painter"
)

type graphService struct{}

func (s *graphService) getGraph(path string) (*dto.Graph, error) {

	err, g := lister.GetGraph(path)
	if err != nil {
		return nil, err
	}

	refd := painter.GetRefdFromHome(g)

	outGraph := dto.Graph{
		Nodes: make(map[string]*dto.Node),
		Edges: make(map[string]*dto.Edge),
	}

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
		outGraph.Nodes[pck.Id()] = &node
	}

	// Collect hits
	for _, pck := range g.Packs {
		if _, ok := refd[pck.Id()]; !ok || !pck.IsHome {
			continue
		}
		node, ok := outGraph.Nodes[pck.Id()]
		if !ok {
			continue
		}
		for _, imp := range pck.Imports() {
			if _, ok := refd[imp.Id()]; !ok {
				continue
			}
			node.OutDeps++
			depNode, ok := outGraph.Nodes[imp.Id()]
			if !ok {
				continue
			}
			depNode.InDeps++
		}
	}

	// Collect edges
	for _, srcPack := range g.Packs {
		if !srcPack.IsHome {
			continue
		}
		for _, imp := range srcPack.Imports() {
			destPack, ok := g.Packs[imp.Id()]
			if !ok {
				continue
			}
			edge := dto.Edge{
				Source: srcPack.Id(),
				Target: imp.Id(),
				Color:  getColor(destPack),
			}
			edgeId := srcPack.Id() + "->" + imp.Id()
			outGraph.Edges[edgeId] = &edge
		}
	}

	return &outGraph, nil
}
