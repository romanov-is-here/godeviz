package godeviz

import (
	"git.dev.cloud.mts.ru/mws/devp/platform-go/pkg/cmdtool"
)

func RootCMD() cmdtool.Command {
	return cmdtool.Command{
		Use:    "godeviz",
		Short:  "godeviz is a command line tool for visualizing package dependency graph",
		Long:   `github.com/romanov-is-here/godeviz/cmd/godeviz`,
		Action: cmdtool.RootAction(),
		SubCommands: []cmdtool.Command{
			ShowCMD(),
			GraphvizCMD(),
		},
	}
}
