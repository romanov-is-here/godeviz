package lister

import "strings"

type PackageInfo struct {
	IsPlatform bool
	IsHome     bool
	IsOuter    bool
	IsStandard bool
	pack       Package
	depGraph   *DepGraph
}

func (p *PackageInfo) Imports() []string {
	return p.pack.Imports
}

func (p *PackageInfo) Name() string {
	if p.IsHome {
		return "🏠" + strings.TrimPrefix(p.id(), p.depGraph.LocalPrefix)
	}
	if p.IsPlatform {
		return "🛠" + strings.TrimPrefix(p.id(), platformPrefix)
	}
	if p.IsStandard {
		return "📦" + p.id()
	}
	if p.IsOuter {
		return "🌐" + p.id()
	}
	return "⚠️" + p.id()
}

func (p *PackageInfo) id() string {
	return p.pack.ImportPath
}

// TODO del
func (p *PackageInfo) Id() string {
	return p.pack.ImportPath
}
