package lister

import (
	"fmt"
	"sort"
	"strings"
)

const (
	platformPrefix = "git.dev.cloud.mts.ru/mws/devp/platform-go"
)

type graph struct {
	home  string
	packs map[string]*PackageInfo
}

func GetGraph(path string) error {
	//ls := golist.DefaultCommandLister()
	ls := &CommandLister{
		Path: path,
	}
	list, err := ls.List()
	if err != nil {
		return err
	}

	mpath, err := ls.ModulePath()
	if err != nil {
		return err
	}

	g := &graph{
		home:  mpath,
		packs: make(map[string]*PackageInfo),
	}

	for _, p := range list {
		//if strings.HasPrefix(p.ImportPath, platformPrefix) {
		//	fmt.Println("🛠 ", p.ImportPath)
		//}
		isPlatform := strings.HasPrefix(p.ImportPath, platformPrefix)
		isHome := strings.HasPrefix(p.ImportPath, mpath)
		isOuter := !isPlatform && !isHome
		g.packs[p.ImportPath] = &PackageInfo{
			IsPlatform: isPlatform,
			IsHome:     isHome,
			IsOuter:    isOuter,
			IsStandard: p.Standard,
			Package:    p,
		}
	}

	//g.print()
	g.hits()

	return nil
}

type hit struct {
	pack *PackageInfo
	from []string
}

func (h *hit) count() int {
	return len(h.from)
}

func (h *hit) register(from string) {
	h.from = append(h.from, from)
}

func (g *graph) hits() {
	hitmap := make(map[string]*hit)

	for _, pi := range g.packs {
		h := &hit{pack: pi}
		hitmap[pi.id()] = h
	}

	for _, pi := range g.packs {
		for _, i := range pi.Package.Imports {
			h, ok := hitmap[i]
			if !ok {
				fmt.Println("HIT NOT FOUND [", i, "] FROM ", pi.name(g.home))
				continue
			}
			h.register(i)
		}
	}

	hsorted := make([]*hit, 0, len(hitmap))

	for _, value := range hitmap {
		hsorted = append(hsorted, value)
	}

	sort.Slice(hsorted, func(i, j int) bool {
		return hsorted[i].count() > hsorted[j].count()
	})

	for _, v := range hsorted {
		if v.count() < 5 {
			continue
		}
		fmt.Println(v.pack.name(g.home), ": ", v.count())
	}

}

func (g *graph) print() {
	for _, pi := range g.packs {
		if !pi.IsHome {
			continue
		}
		fmt.Println(pi.name(g.home))
		for _, i := range pi.Package.Imports {
			f, ok := g.packs[i]
			if f.IsStandard || f.IsOuter {
				continue
			}
			fmt.Print("\t")
			if !ok {
				fmt.Println("⚠️⚠️" + i)
				continue
			}
			fmt.Println(f.name(g.home))
		}
	}
}

func (p *PackageInfo) name(home string) string {
	if p.IsHome {
		return "🏠" + strings.TrimPrefix(p.id(), home)
	}
	if p.IsPlatform {
		return "🛠" + strings.TrimPrefix(p.id(), platformPrefix)
	}
	if p.IsOuter {
		return "🌐" + p.id()
	}
	return "⚠️" + p.id()
}

func (p *PackageInfo) id() string {
	return p.Package.ImportPath
}

type PackageInfo struct {
	IsPlatform bool
	IsHome     bool
	IsOuter    bool
	IsStandard bool
	Package    Package
}
