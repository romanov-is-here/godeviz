package graphcontroller

import (
	"strconv"

	"github.com/romanov-is-here/godeviz/internal/api/controllers/graphcontroller/dto"
	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
	"github.com/romanov-is-here/godeviz/internal/graph/lister"
)

type graphService struct{}

func (s *graphService) getGraph(path string, filter *depgraph.Filter) (*dto.Graph, error) {

	err, g := lister.GetGraph(path, filter)
	if err != nil {
		return nil, err
	}

	outGraph := dto.Graph{
		Nodes: make(map[string]*dto.Node),
		Edges: make(map[string]*dto.Edge),
	}

	// Collect nodes
	for _, pck := range g.Packs {
		imports := make([]*dto.Import, 0)
		for i, imp := range pck.Imports {
			dtoImp := &dto.Import{
				Id:   strconv.Itoa(i),
				Name: imp.Id,
			}
			imports = append(imports, dtoImp)
		}
		node := dto.Node{
			Name:        pck.Name,
			IsHome:      pck.IsHome,
			Color:       getColor(pck),
			FanInCount:  pck.FanInCount,
			FanOutCount: pck.FanOutCount,
			Imports:     imports,
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
