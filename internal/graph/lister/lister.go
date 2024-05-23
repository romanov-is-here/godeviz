package lister

import (
	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
)

func GetGraph(path string, filter *depgraph.Filter) (error, *depgraph.DepGraph) {
	ls := &CommandLister{
		Path: path,
	}
	list, err := ls.List()
	if err != nil {
		return err, nil
	}

	mpath, err := ls.ModulePath()
	if err != nil {
		return err, nil
	}

	b := depgraph.NewBuilder(mpath)

	for _, p := range list {
		b.Add(p)
	}

	b.SetFilter(filter)

	return nil, b.Build()
}
