package lister

import (
	"fmt"
	"sort"
)

type DepGraph struct {
	LocalPrefix string
	Packs       map[string]*PackageInfo
}

func (g *DepGraph) add(pack *PackageInfo) {
	pack.depGraph = g
	g.Packs[pack.id()] = pack
}

func (g *DepGraph) Hits() Hits {
	hitmap := make(map[string]*Hit)

	for _, pi := range g.Packs {
		h := &Hit{pack: pi}
		hitmap[pi.id()] = h
	}

	for _, pi := range g.Packs {
		for _, i := range pi.pack.Imports {
			h, ok := hitmap[i]
			if !ok {
				fmt.Println("HIT NOT FOUND [", i, "] FROM ", pi.Name())
				continue
			}
			h.add(i)
		}
	}

	hsorted := make([]*Hit, 0, len(hitmap))

	for _, value := range hitmap {
		hsorted = append(hsorted, value)
	}

	sort.Slice(hsorted, func(i, j int) bool {
		return hsorted[i].Count() > hsorted[j].Count()
	})

	return Hits{
		Hits: hsorted,
	}
}

func (g *DepGraph) print() {
	for _, pi := range g.Packs {
		if !pi.IsHome {
			continue
		}
		fmt.Println(pi.Name())
		for _, i := range pi.pack.Imports {
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
