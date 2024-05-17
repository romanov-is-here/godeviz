package main

import (
	"fmt"

	"github.com/romanov-is-here/godeviz/internal/lister"
)

func main() {
	err := lister.GetGraph("/Users/ilyaromanov/Projects/api")
	//err := lister.GetGraph("/Users/ilyaromanov/Projects/platform-go/cmd/devp")
	//err := lister.GetGraph("./...")
	if err != nil {
		fmt.Println("error happened: ", err)
	}
}
