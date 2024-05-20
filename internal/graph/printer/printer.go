package printer

import (
	"fmt"

	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
)

func Graph(g *depgraph.DepGraph) {
	for _, pi := range g.Packs {
		if !pi.IsHome {
			continue
		}
		fmt.Println(pi.Name())
		for _, imp := range pi.Imports() {
			f, ok := g.Packs[imp.Id()]
			if f.IsStandard || f.IsOuter {
				//continue
			}
			fmt.Print("\t")
			if !ok {
				fmt.Println("⚠️⚠️" + imp.Id())
				continue
			}
			fmt.Println(f.Name())
		}
	}
}

func Hits(g *depgraph.DepGraph) {
	hsorted := g.Hits()
	for _, h := range hsorted.Hits {
		if h.Count() < 5 {
			continue
		}
		fmt.Println(h.Name(), ": ", h.Count())
	}
}
