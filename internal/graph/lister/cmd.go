package lister

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/romanov-is-here/godeviz/internal/graph/gomodel"
)

type CommandLister struct {
	Path string
}

// List returns the struct unmarshalled from `go list --json` call to be used for example as an input of project formatter
func (c *CommandLister) List() ([]gomodel.Package, error) {
	return listPackages(c.run, c.Path)
}

// Modules returns the struct unmarshalled from `go list -m --json`
func (c *CommandLister) Modules() ([]gomodel.Module, error) {
	return listModules(c.run)
}

// ModulePath returns main module path or an empty string if the module is not found
func (c *CommandLister) ModulePath() (string, error) {
	return getModulePath(c.run)
}

func getModulePath(run runFunc) (string, error) {
	modules, err := listModules(run)
	if err != nil {
		return "", err
	}

	for _, m := range modules {
		if m.Main && m.Path != "command-line-arguments" {
			return m.Path, nil
		}
	}

	return "", nil
}

func (c *CommandLister) run(out *bytes.Buffer, arg ...string) error {
	cmd := exec.Command("go", arg...) // #nosec G204
	cmd.Dir = c.Path
	errBuf := &bytes.Buffer{}
	cmd.Stderr = errBuf
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", err, errBuf.String())
	}
	return nil
}

func listPackages(run runFunc, path string) ([]gomodel.Package, error) {
	var x []gomodel.Package
	//data, err := execute(run, "list", "--json=ImportPath,Imports,Standard", "--deps", "all", path)
	//data, err := execute(run, "list", "--json=ImportPath,Imports,Standard", "--deps", path)
	data, err := execute(run, "list", "--json=ImportPath,Imports,Standard", "--deps", "./...")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}
	return x, nil
}

func listModules(run runFunc) ([]gomodel.Module, error) {
	var x []gomodel.Module
	data, err := execute(run, "list", "-m", "--json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}

	return x, nil
}

type runFunc func(out *bytes.Buffer, arg ...string) error

var fixReg = regexp.MustCompile(`}\s+{`)
var replaceDst = []byte(`},{`)

func execute(run runFunc, arg ...string) ([]byte, error) {
	outBuf := &bytes.Buffer{}
	outBuf.WriteByte('[')
	err := run(outBuf, arg...)
	if err != nil {
		return nil, err
	}
	outBuf.WriteByte(']')
	return fixReg.ReplaceAll(outBuf.Bytes(), replaceDst), nil
}
