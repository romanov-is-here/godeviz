package painter

import (
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"

	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
)

func Graphviz(gr *depgraph.DepGraph, output string) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if errd := graph.Close(); errd != nil {
			log.Fatal(errd)
		}
		g.Close()
	}()

	nodes := make(map[string]*cgraph.Node)

	for k, pck := range gr.Packs {
		node, errn := graph.CreateNode(pck.Id)
		node.SetColor(getColor(pck))
		if errn != nil {
			return errn
		}
		nodes[k] = node
	}

	for _, srcPack := range gr.Packs {
		srcNode, _ := nodes[srcPack.Id]
		for _, imp := range srcPack.Imports {
			destPack, ok := gr.Packs[imp.Id]
			if !ok {
				continue // TODO is it possible?
			}
			destNode, ok := nodes[destPack.Id]
			if !ok {
				continue // TODO is it possible?
			}
			edge, ederr := graph.CreateEdge("", srcNode, destNode)
			edge.SetColor(getColor(destPack))
			if ederr != nil {
				return ederr
			}
		}
	}

	if err2 := g.RenderFilename(graph, graphviz.SVG, output); err2 != nil {
		log.Fatal(err2)
	}
	return nil
}

func getColor(p *depgraph.Package) string {
	if p.IsHome {
		return "brown"
	}
	if p.IsPlatform {
		return "purple"
	}
	if p.IsStandard {
		return "green"
	}
	if p.IsOuter {
		return "lightblue"
	}
	return "lightgrey"
}
