package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gorilla/mux"

	"github.com/romanov-is-here/godeviz/internal/api"
	"github.com/romanov-is-here/godeviz/internal/vueapp"
)

func main() {
	r := mux.NewRouter()

	api.Setup(r)
	vueapp.Setup(r)

	go run()
	go waitAndShow()

	waitForStop()
}

func waitForStop() {
	log.Println("Godeviz is running. Press Enter to stop.")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("%w", err)
	}
	log.Println("Godeviz stopped")
}

func waitAndShow() {
	time.Sleep(time.Second)
	err := openBrowser("http://localhost:8080")
	if err != nil {
		log.Println("Open your browser at http://localhost:8080")
	}
}

func run() {
	log.Println("Starting godeviz on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func openBrowser(url string) error {
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
