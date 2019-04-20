// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ghCommand = []string{"run", "../main.go"}
)

func TestGH(t *testing.T) {
	fmt.Println("TestGH")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// no flags set
	currentCmd := ghCommand
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.Contains(stdout, "A collection of git helper commands to make my life a little easier")
	assert.Contains(stdout, "--version         version for gh")
	outbuf.Reset()
	errbuf.Reset()
}

func TestGHWithVersion(t *testing.T) {
	fmt.Println("TestGHWithVersion")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// version flag set
	currentCmd := append(ghCommand, "--version")
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.Contains(stdout, "You're running gh version 2.0.0")
	outbuf.Reset()
	errbuf.Reset()
}

func TestGHWithHelp(t *testing.T) {
	fmt.Println("TestGHWithHelp")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// help flag set
	currentCmd := append(ghCommand, "--help")
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.Contains(stdout, "A collection of git helper commands to make my life a little easier")
	assert.Contains(stdout, "Use \"gh [command] --help\" for more information about a command.")
	outbuf.Reset()
	errbuf.Reset()
}
