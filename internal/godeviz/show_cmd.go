package godeviz

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"git.dev.cloud.mts.ru/mws/devp/platform-go/pkg/cmdtool"
	"git.dev.cloud.mts.ru/mws/devp/platform-go/pkg/zapctx"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/romanov-is-here/godeviz/internal/api"
	"github.com/romanov-is-here/godeviz/internal/vueapp"
)

func ShowCMD() cmdtool.Command {
	return cmdtool.Command{
		Use:   "show",
		Short: "show project dependency graph",

		Action: cmdtool.RunAction(&showMeRunner{}),
	}
}

type showMeRunner struct {
	port int
}

func (r *showMeRunner) BindFlags(fs cmdtool.FlagSet) {
	fs.IntVar(&r.port, "port", 8080, "port to run http listener on")
}

func (r *showMeRunner) Run(ctx context.Context, _ []string) error {
	router := mux.NewRouter()

	api.Setup(router)
	vueapp.Setup(router)

	go r.run()
	go r.waitAndShow(ctx)

	r.waitForStop(ctx)

	return nil
}

func (r *showMeRunner) waitForStop(ctx context.Context) {
	zapctx.Info(ctx, "Godeviz is running. Press Enter to stop.", zap.Int("port", r.port))
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		// doesn't matter
	}
	zapctx.Info(ctx, "Godeviz stopped")
}

func (r *showMeRunner) waitAndShow(ctx context.Context) {
	addr := fmt.Sprintf("http://localhost:%d", r.port)
	r.waitForPingAPI()
	err := open(addr)
	if err != nil {
		zapctx.Info(ctx, "Open your browser", zap.String("address", addr))
	}
}

func (r *showMeRunner) run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", r.port), nil))
}

func (r *showMeRunner) waitForPingAPI() {
	url := fmt.Sprintf("%s/api/ping", r.address())
	for {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}

		if resp.StatusCode == http.StatusOK {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (r *showMeRunner) address() string {
	return fmt.Sprintf("http://localhost:%d", r.port)
}
