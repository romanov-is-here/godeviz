package depgraph

type Package struct {
	Name       string
	IsPlatform bool
	IsHome     bool
	IsOuter    bool
	IsStandard bool
	Id         string
	Imports    []*Import
	FanIn      int
	FanOut     int
}

type DepGraph struct {
	Packs map[string]*Package
}

type Import struct {
	Id string
}

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
