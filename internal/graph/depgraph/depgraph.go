package depgraph

import (
	"fmt"
	"sort"
	"strings"

	"github.com/romanov-is-here/godeviz/internal/graph/gomodel"
)

const (
	platformPrefix = "git.dev.cloud.mts.ru/mws/devp/platform-go"
)

func NewDepGraph(localPrefix string) *DepGraph {
	return &DepGraph{
		localPrefix: localPrefix,
		Packs:       make(map[string]*PackageInfo),
	}
}

type DepGraph struct {
	localPrefix string
	Packs       map[string]*PackageInfo
}

func (g *DepGraph) Add(p gomodel.Package) {
	isPlatform := strings.HasPrefix(p.ImportPath, platformPrefix)
	isHome := strings.HasPrefix(p.ImportPath, g.localPrefix)
	isOuter := !isPlatform && !isHome

	imports := make([]*Import, 0)
	for _, i := range p.Imports {
		imports = append(imports, &Import{id: i})
	}

	pack := &PackageInfo{
		IsPlatform: isPlatform,
		IsHome:     isHome,
		IsOuter:    isOuter,
		IsStandard: p.Standard,
		depGraph:   g,
		id:         p.ImportPath,
		imports:    imports,
	}

	g.Packs[pack.Id()] = pack
}

func (g *DepGraph) Hits() Hits {
	hitmap := make(map[string]*Hit)

	for _, pi := range g.Packs {
		h := &Hit{pack: pi}
		hitmap[pi.Id()] = h
	}

	for _, pi := range g.Packs {
		for _, imp := range pi.Imports() {
			h, ok := hitmap[imp.Id()]
			if !ok {
				fmt.Println("HIT NOT FOUND [", imp.Id(), "] FROM ", pi.Name())
				continue
			}
			h.add(imp.Id())
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
