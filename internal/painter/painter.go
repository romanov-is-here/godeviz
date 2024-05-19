package painter

import (
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"

	"github.com/romanov-is-here/godeviz/internal/lister"
)

func Graphviz(gr *lister.DepGraph) error {
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

	refd := getRefdFromHome(gr)

	nodes := make(map[string]*cgraph.Node)

	for k, pck := range gr.Packs {
		if _, ok := refd[pck.Id()]; !ok {
			continue
		}
		node, errn := graph.CreateNode(pck.Id())
		node.SetColor(getColor(pck))
		if errn != nil {
			return errn
		}
		nodes[k] = node
	}
	for _, srcPack := range gr.Packs {
		if !srcPack.IsHome {
			continue
		}
		srcNode, _ := nodes[srcPack.Id()]
		for _, imp := range srcPack.Imports() {
			destPack, ok := gr.Packs[imp]
			if !ok {
				continue // TODO is it possible?
			}
			destNode, ok := nodes[destPack.Id()]
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

	if err := g.RenderFilename(graph, graphviz.SVG, "./graph.svg"); err != nil {
		log.Fatal(err)
	}
	return nil
}

func getRefdFromHome(gr *lister.DepGraph) map[string]bool {
	refd := make(map[string]bool)

	for _, v := range gr.Packs {
		if v.IsHome {
			refd[v.Id()] = true
			for _, imp := range v.Imports() {
				refd[imp] = true
			}
		}
	}
	return refd
}

func getColor(p *lister.PackageInfo) string {
	if p.IsHome {
		return "black"
	}
	if p.IsPlatform {
		return "green"
	}
	if p.IsStandard {
		return "green"
	}
	if p.IsOuter {
		return "red"
	}
	return "lightgrey"
}
