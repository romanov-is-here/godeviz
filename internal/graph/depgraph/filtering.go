package depgraph

func (b *Builder) applyFilter() {
	var importsFromHome map[string]bool

	if b.filter.OnlyHomeRoots {
		importsFromHome = b.getImportsFromHome()
	}

	// Filtering packs
	focusPackage := b.packs[b.filter.FocusPackage]
	for k, p := range b.packs {
		if !b.passPackageFilter(p, importsFromHome, focusPackage) {
			delete(b.packs, k)
		}
	}

	// Filtering imports
	for _, p := range b.packs {
		filtered := make([]*packageImport, 0)
		for _, imp := range p.imports {
			if b.passImportFilter(p, imp, focusPackage) {
				filtered = append(filtered, imp)
			}

			p.imports = filtered
		}
	}
}

func (b *Builder) passImportFilter(p *packageInfo, imp *packageImport, focusPackage *packageInfo) bool {
	f := b.filter

	_, ok := b.packs[imp.id]
	if !ok {
		return false
	}

	if f.OnlyHomeRoots && !p.isHome {
		return false
	}

	if !b.passFocusFilter(imp.id, focusPackage) {
		return false
	}

	return true
}

func (b *Builder) passPackageFilter(p *packageInfo, importsFromHome map[string]bool, focusPackage *packageInfo) bool {
	f := b.filter

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

	if f.OnlyHomeRoots {
		_, ok := importsFromHome[p.id]
		if !ok {
			return false
		}
	}

	if !b.passFocusFilter(p.id, focusPackage) {
		return false
	}

	return true
}

func (b *Builder) passFocusFilter(pid string, focusPackage *packageInfo) bool {
	f := b.filter

	if f.FocusPackage == "" || f.FocusPackage == pid {
		return true
	}

	if focusPackage == nil {
		return false // TODO можно возвращать ошибку, но не хочется. Если не нашли - ничего не покажем
	}

	if f.FocusType == OutsOnly || f.FocusType == Both {
		return focusPackage.hasImport(pid)
	}

	p, ok := b.packs[pid]
	if !ok {
		return false
	}
	if f.FocusType == InsOnly || f.FocusType == Both {
		return p.hasImport(f.FocusPackage)
	}

	return true // should be unreachable
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
