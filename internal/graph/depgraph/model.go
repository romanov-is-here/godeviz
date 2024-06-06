package depgraph

type Package struct {
	Name        string
	IsPlatform  bool
	IsHome      bool
	IsOuter     bool
	IsStandard  bool
	Id          string
	Imports     []*Import
	FanInCount  int
	FanOutCount int
}

type DepGraph struct {
	Packs      map[string]*Package
	ModuleName string
}

type Import struct {
	Id string
}

type Filter struct {
	ShowHome      bool // TODO does false makes sense?
	ShowStandard  bool
	ShowPlatform  bool
	ShowOuter     bool
	OnlyHomeRoots bool // TODO does false makes sense?
	FocusPackage  string
	FocusType     FocusType
	ShowRatio     bool
}

type FocusType int

const (
	OutsOnly FocusType = iota
	InsOnly
	Both
)

func NewDefaultFilter() *Filter {
	return &Filter{
		ShowHome:      true,
		ShowPlatform:  true,
		ShowOuter:     true,
		OnlyHomeRoots: true,
	}
}
