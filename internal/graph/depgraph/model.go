package depgraph

type Package struct {
	Name       string
	IsPlatform bool
	IsHome     bool
	IsOuter    bool
	IsStandard bool
	Id         string
	Imports    []*Import
}

type DepGraph struct {
	Packs map[string]*Package
}

type Import struct {
	Id string
}
