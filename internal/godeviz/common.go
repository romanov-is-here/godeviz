package godeviz

import (
	"fmt"
	"os/exec"
	"runtime"
)

func open(arg string) error {
	var cmd string
	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
	case "windows":
		cmd = "cmd /c start"
	case "darwin":
		cmd = "open"
	default:
		return fmt.Errorf("cannot open '%s' on %s", arg, runtime.GOOS)
	}
	return exec.Command(cmd, arg).Start()
}
