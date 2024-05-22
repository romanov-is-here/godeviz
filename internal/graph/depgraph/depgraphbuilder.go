package depgraph

import (
	"strings"

	"github.com/romanov-is-here/godeviz/internal/graph/gomodel"
)

const (
	platformPrefix = "git.dev.cloud.mts.ru/mws/devp/platform-go"
)

type Builder struct {
	localPrefix string
	packs       map[string]*packageInfo
	filter      *Filter
}

func NewBuilder(localPrefix string) *Builder {
	return &Builder{
		localPrefix: localPrefix,
		packs:       make(map[string]*packageInfo),
		filter:      NewDefaultFilter(),
	}
}

func (b *Builder) Add(p gomodel.Package) {
	isPlatform := strings.HasPrefix(p.ImportPath, platformPrefix)
	isHome := strings.HasPrefix(p.ImportPath, b.localPrefix)
	isOuter := !isPlatform && !isHome

	imports := make([]*packageImport, 0)
	for _, i := range p.Imports {
		imports = append(imports, &packageImport{id: i})
	}

	pack := &packageInfo{
		isPlatform: isPlatform,
		isHome:     isHome,
		isOuter:    isOuter,
		isStandard: p.Standard,
		builder:    b,
		id:         p.ImportPath,
		imports:    imports,
	}

	b.packs[pack.id] = pack
}

func (b *Builder) SetFilter(filter *Filter) {
	if filter == nil {
		filter = NewDefaultFilter()
	}
	b.filter = filter
}

func (b *Builder) Build() *DepGraph {
	b.applyFilter()
	b.collectHits()
	g := &DepGraph{
		Packs: make(map[string]*Package),
	}

	for k, v := range b.packs {
		g.Packs[k] = newPackage(v)
	}

	return g
}

func (b *Builder) collectHits() {
	for _, pck := range b.packs {
		pck.fanOut = len(pck.imports)
		for _, imp := range pck.imports {
			if destPack, ok := b.packs[imp.id]; ok {
				destPack.fanOut++
			}
		}
	}
}

type packageInfo struct {
	isPlatform bool
	isHome     bool
	isOuter    bool
	isStandard bool
	id         string
	imports    []*packageImport
	builder    *Builder
	fanIn      int
	fanOut     int
}

func newPackage(p *packageInfo) *Package {
	imports := make([]*Import, 0)
	for _, imp := range p.imports {
		imports = append(imports, &Import{Id: imp.id})
	}
	return &Package{
		Name:       p.Name(),
		IsPlatform: p.isPlatform,
		IsHome:     p.isHome,
		IsOuter:    p.isOuter,
		IsStandard: p.isStandard,
		Id:         p.id,
		Imports:    imports,
		FanIn:      p.fanIn,
		FanOut:     p.fanOut,
	}
}

func (p *packageInfo) Name() string {
	if p.isHome {
		return "🏠" + strings.TrimPrefix(p.id, p.builder.localPrefix)
	}
	if p.isPlatform {
		return "🛠" + strings.TrimPrefix(p.id, platformPrefix)
	}
	if p.isStandard {
		return "📦" + p.id
	}
	if p.isOuter {
		return "🌐" + p.id
	}
	return "⚠️" + p.id // this is unexpected
}

type packageImport struct {
	id string
}
