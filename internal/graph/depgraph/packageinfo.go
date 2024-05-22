package depgraph

import (
	"strings"
)

type PackageInfo struct {
	IsPlatform bool
	IsHome     bool
	IsOuter    bool
	IsStandard bool
	id         string
	imports    []*Import
	depGraph   *DepGraph
}

func (p *PackageInfo) Imports() []*Import {
	return p.imports
}

func (p *PackageInfo) Name() string {
	if p.IsHome {
		return "🏠" + strings.TrimPrefix(p.id, p.depGraph.localPrefix)
	}
	if p.IsPlatform {
		return "🛠" + strings.TrimPrefix(p.id, platformPrefix)
	}
	if p.IsStandard {
		return "📦" + p.id
	}
	if p.IsOuter {
		return "🌐" + p.id
	}
	return "⚠️" + p.id
}

func (p *PackageInfo) Id() string {
	return p.id
}
