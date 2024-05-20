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
		return "🏠" + strings.TrimPrefix(p.Id(), p.depGraph.localPrefix)
	}
	if p.IsPlatform {
		return "🛠" + strings.TrimPrefix(p.Id(), platformPrefix)
	}
	if p.IsStandard {
		return "📦" + p.Id()
	}
	if p.IsOuter {
		return "🌐" + p.Id()
	}
	return "⚠️" + p.Id()
}

func (p *PackageInfo) Id() string {
	return p.id
}
