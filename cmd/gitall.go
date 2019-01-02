// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/retgits/gh/util"

	"github.com/spf13/cobra"
)

// gitAllCmd represents the all command
var gitAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Stage all unstaged files",
	Run:   runGitAll,
}

// Flags
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitAllCmd)
}

// runGitAll is the actual execution of the command
func runGitAll(cmd *cobra.Command, args []string) {
	currentDirectory, err := util.GetCurrentDirectory()
	if err != nil {
		fmt.Printf("An error occurred while resolving current directory: %s", err.Error())
		os.Exit(2)
	}
	cmdExec := exec.Command("sh", "-c", "git add -A")
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	fmt.Println(currentDirectory)
	cmdExec.Run()
}
