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
}

type packageInfo struct {
	isPlatform bool
	isHome     bool
	isOuter    bool
	isStandard bool
	id         string
	imports    []*packageImport
	builder    *Builder
}

type packageImport struct {
	id string
}

func NewBuilder(localPrefix string) *Builder {
	return &Builder{
		localPrefix: localPrefix,
		packs:       make(map[string]*packageInfo),
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

func (b *Builder) Build() *DepGraph {
	// TODO filter
	// TODO hits
	g := &DepGraph{
		Packs: make(map[string]*Package),
	}

	for k, v := range b.packs {
		g.Packs[k] = newPackage(v)
	}

	return g
}

func newPackage(v *packageInfo) *Package {
	imports := make([]*Import, 0)
	for _, imp := range v.imports {
		imports = append(imports, &Import{Id: imp.id})
	}
	return &Package{
		Name:       v.Name(),
		IsPlatform: v.isPlatform,
		IsHome:     v.isHome,
		IsOuter:    v.isOuter,
		IsStandard: v.isStandard,
		Id:         v.id,
		Imports:    imports,
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
