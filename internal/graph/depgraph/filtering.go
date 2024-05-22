package depgraph

func (b *Builder) applyFilter() {
	var importsFromHome map[string]bool

	if !b.filter.ShowNonHomeRoots {
		importsFromHome = b.getImportsFromHome()
	}

	for k, p := range b.packs {
		if !b.filter.passFilter(p, importsFromHome) {
			delete(b.packs, k)
		}
	}
}

func (f *Filter) passFilter(p *packageInfo, importsFromHome map[string]bool) bool {
	if !f.ShowHome && p.isHome {
		return false
	}

	if !f.ShowPlatform && p.isPlatform {
		return false
	}

	if !f.ShowOuter && p.isOuter {
		return false
	}

	if !f.ShowStandard && p.isStandard {
		return false
	}

	if !f.ShowNonHomeRoots {
		_, ok := importsFromHome[p.id]
		if !ok {
			return false
		}
	}

	return true
}

func (b *Builder) getImportsFromHome() map[string]bool {
	importsFromHome := make(map[string]bool)

	for _, v := range b.packs {
		if v.isHome {
			importsFromHome[v.id] = true
			for _, imp := range v.imports {
				importsFromHome[imp.id] = true
			}
		}
	}
	return importsFromHome
}
