package godeviz

import (
	"context"
	"fmt"

	"git.dev.cloud.mts.ru/mws/devp/platform-go/pkg/cmdtool"
	"git.dev.cloud.mts.ru/mws/devp/platform-go/pkg/zapctx"
	"go.uber.org/zap"

	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
	"github.com/romanov-is-here/godeviz/internal/graph/lister"
	"github.com/romanov-is-here/godeviz/internal/graph/painter"
)

func GraphvizCMD() cmdtool.Command {
	return cmdtool.Command{
		Use:     "graphviz",
		Aliases: []string{"gviz", "svg"},
		Short:   "show project dependency graph in graphviz",

		Action: cmdtool.RunAction(&graphvizRunner{}),
	}
}

type graphvizRunner struct {
	output string
	path   string
}

func (r *graphvizRunner) BindFlags(fs cmdtool.FlagSet) {
	fs.StringVar(&r.output, "output", "graph.svg", "output file path")
	fs.StringVar(&r.path, "path", ".", "path to the project")
}

func (r *graphvizRunner) Run(ctx context.Context, _ []string) error {
	err, g := lister.GetGraph(r.path, depgraph.NewDefaultFilter())
	if err != nil {
		return fmt.Errorf("failed to make a graph: %w", err)
	}

	if err = painter.Graphviz(g, r.output); err != nil {
		return fmt.Errorf("failed to paint a graph: %w", err)
	}

	_ = open(r.output)
	zapctx.Info(ctx, "graph is ready", zap.String("file", r.output))

	return nil
}
