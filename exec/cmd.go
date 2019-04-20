// Package exec is a wrapper around the os/exec package
package exec

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	macOSExecutable   = "sh"
	macOSFlag         = "-c"
	linuxExecutable   = "sh"
	linuxFlag         = "-c"
	windowsExecutable = "cmd.exe"
	windowsFlag       = "/c"
)

// RunCmd runs a specific command in the default shell of the operating system.
func RunCmd(args string) error {
	var executable string
	var flag string

	switch strings.ToLower(runtime.GOOS) {
	case "darwin":
		executable = macOSExecutable
		flag = macOSFlag
	case "windows":
		executable = windowsExecutable
		flag = windowsFlag
	case "linux":
		executable = linuxExecutable
		flag = linuxFlag
	}

	cmd := exec.Command(executable, flag, args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	dir, err := currentDirectory()
	if err != nil {
		return err
	}
	cmd.Dir = dir
	return cmd.Run()
}

func currentDirectory() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
