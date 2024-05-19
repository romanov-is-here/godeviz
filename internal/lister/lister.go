package lister

import (
	"strings"
)

const (
	platformPrefix = "git.dev.cloud.mts.ru/mws/devp/platform-go"
)

func GetGraph(path string) (error, *DepGraph) {
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

	g := &DepGraph{
		LocalPrefix: mpath,
		Packs:       make(map[string]*PackageInfo),
	}

	for _, p := range list {
		isPlatform := strings.HasPrefix(p.ImportPath, platformPrefix)
		isHome := strings.HasPrefix(p.ImportPath, mpath)
		isOuter := !isPlatform && !isHome
		pack := &PackageInfo{
			IsPlatform: isPlatform,
			IsHome:     isHome,
			IsOuter:    isOuter,
			IsStandard: p.Standard,
			pack:       p,
		}
		g.add(pack)
	}

	return nil, g
}
