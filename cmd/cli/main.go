package main

import (
	"fmt"

	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
	"github.com/romanov-is-here/godeviz/internal/graph/lister"
	"github.com/romanov-is-here/godeviz/internal/graph/painter"
)

func main() {
	//err, g := lister.GetGraph("/Users/ilyaromanov/Projects/api", depgraph.NewDefaultFilter())
	err, g := lister.GetGraph("/Users/ilyaromanov/Projects/godeviz", depgraph.NewDefaultFilter())
	//err, g := lister.GetGraph("/Users/ilyaromanov/Projects/platform-go/cmd/devp", depgraph.NewDefaultFilter())
	//err := lister.GetGraph("./...")
	if err != nil {
		fmt.Println("error happened: ", err)
	}
	//fmt.Println(printer.DotNotation(g))
	err = painter.Graphviz(g)
	if err != nil {
		fmt.Println("error painting: ", err)
	}
}
