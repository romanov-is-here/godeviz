package main

import (
	"git.dev.cloud.mts.ru/mws/devp/platform-go/pkg/cmdtool/cobra"

	"github.com/romanov-is-here/godeviz/internal/godeviz"
)

func main() {
	cobra.Main(godeviz.RootCMD())
}
