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
	cloneCommand = []string{"run", "../main.go", "clone"}
	tmpCloneRepo = "https://github.com/retgits/grpcrest-proxy"
)

func TestClone(t *testing.T) {
	fmt.Println("TestClone")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// no flags set
	currentCmd := cloneCommand
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	// Unset the environment variable in case it was set
	os.Unsetenv("GITBASEFOLDER")

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.Contains(stdout, "Error: There was no URL provided")
	outbuf.Reset()
	errbuf.Reset()
}

func TestCloneWithBaseFlag(t *testing.T) {
	fmt.Println("TestCloneWithBaseFlag")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// base flags set
	currentCmd := append(cloneCommand, "--base", os.Getenv("TESTDIR"))
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	// Unset the environment variable in case it was set
	os.Unsetenv("GITBASEFOLDER")

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.Contains(stdout, "Error: There was no URL provided")
	outbuf.Reset()
	errbuf.Reset()
}

func TestCloneWithBaseFlagAndURL(t *testing.T) {
	fmt.Println("TestCloneWithBaseFlagAndURL")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// base flags set
	currentCmd := append(cloneCommand, "--base", os.Getenv("TESTDIR"), tmpCloneRepo)
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	// Unset the environment variable in case it was set
	os.Unsetenv("GITBASEFOLDER")

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.NotContains(stdout, "gh clone is a simple git clone command")
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), "github.com", "retgits", "grpcrest-proxy", "LICENSE"))
	outbuf.Reset()
	errbuf.Reset()
	os.RemoveAll(filepath.Join(os.Getenv("TESTDIR"), "github.com"))
}

func TestCloneWithBaseFolderAndURL(t *testing.T) {
	fmt.Println("TestCloneWithBaseFolderAndURL")
	assert := assert.New(t)

	var outbuf, errbuf bytes.Buffer

	// base flags set
	currentCmd := append(cloneCommand, tmpCloneRepo)
	cmd := exec.Command("go", currentCmd...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	// Update or set the environment variable in case it was set
	os.Setenv("GITBASEFOLDER", os.Getenv("TESTDIR"))

	err := cmd.Run()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") {
		fmt.Println(err.Error())
	}
	stdout := outbuf.String()
	assert.NotContains(stdout, "gh clone is a simple git clone command")
	assert.FileExists(filepath.Join(os.Getenv("TESTDIR"), "github.com", "retgits", "grpcrest-proxy", "LICENSE"))
	outbuf.Reset()
	errbuf.Reset()
	os.RemoveAll(filepath.Join(os.Getenv("TESTDIR"), "github.com"))
}
