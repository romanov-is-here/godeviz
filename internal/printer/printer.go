package printer

import (
	"fmt"

	"github.com/romanov-is-here/godeviz/internal/lister"
)

func Graph(g *lister.DepGraph) {
	for _, pi := range g.Packs {
		if !pi.IsHome {
			continue
		}
		fmt.Println(pi.Name())
		for _, i := range pi.Imports() {
			f, ok := g.Packs[i]
			if f.IsStandard || f.IsOuter {
				//continue
			}
			fmt.Print("\t")
			if !ok {
				fmt.Println("⚠️⚠️" + i)
				continue
			}
			fmt.Println(f.Name())
		}
	}
}

func Hits(g *lister.DepGraph) {
	hsorted := g.Hits()
	for _, h := range hsorted.Hits {
		if h.Count() < 5 {
			continue
		}
		fmt.Println(h.Name(), ": ", h.Count())
	}
}
