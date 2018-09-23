// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lambdaCommand = []string{"run", "../main.go", "lambda"}
	tmpLambdaName = "myLambda"
)

func TestLambda(t *testing.T) {
	fmt.Println("TestLambda")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// no flags set
	currentCmd := lambdaCommand
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.Contains(stdout, "required flag(s) \"name\" not set")
	outbuf.Reset()
	errbuf.Reset()
}

func TestLambdaWithNameAndBase(t *testing.T) {
	fmt.Println("TestLambdaWithNameAndBase")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// name flags set
	currentCmd := append(lambdaCommand, "--name", tmpLambdaName, "--base", os.Getenv("TESTDIR"))
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.NotContains(stdout, "Usage:")
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), tmpLambdaName, ".gitignore"))
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), tmpLambdaName, "LICENSE"))
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), tmpLambdaName, "main.go"))
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), tmpLambdaName, "makefile"))
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), tmpLambdaName, "template.yml"))
	outbuf.Reset()
	errbuf.Reset()
	os.RemoveAll(filepath.Join(os.Getenv("TESTDIR"), tmpLambdaName))
}
