package depgraph

type Filter struct {
	ShowHome         bool
	ShowStandard     bool
	ShowPlatform     bool
	ShowOuter        bool
	ShowNonHomeRoots bool
}

func NewDefaultFilter() *Filter {
	return &Filter{
		ShowHome:         true,
		ShowStandard:     false,
		ShowPlatform:     true,
		ShowOuter:        true,
		ShowNonHomeRoots: false,
	}
}
