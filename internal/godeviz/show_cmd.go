package godeviz

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
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
	r.waitForPingAPI()
	err := r.openBrowser("http://localhost:8080")
	if err != nil {
		zapctx.Info(ctx, "Open your browser at http://localhost:8080")
	}
}

func (r *showMeRunner) run() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (r *showMeRunner) openBrowser(url string) error {
	var cmd string
	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
	case "windows":
		cmd = "cmd /c start"
	case "darwin":
		cmd = "open"
	default:
		return fmt.Errorf("cannot open browser on %s", runtime.GOOS)
	}
	return exec.Command(cmd, url).Start()
}

func (r *showMeRunner) waitForPingAPI() {
	url := "http://localhost:8080/api/ping"
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
