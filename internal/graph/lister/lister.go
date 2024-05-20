package lister

import (
	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
)

func GetGraph(path string) (error, *depgraph.DepGraph) {
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

	g := depgraph.NewDepGraph(mpath)

	for _, p := range list {
		g.Add(p)
	}

	return nil, g
}
