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

// gitAmendCmd represents the git command
var gitAmendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Use the last commit message and amend your stuffs",
	Run:   runGitAmend,
}

// Flags
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitAmendCmd)
}

// runGitAmend is the actual execution of the command
func runGitAmend(cmd *cobra.Command, args []string) {
	currentDirectory, err := util.GetCurrentDirectory()
	if err != nil {
		fmt.Printf("An error occured while resolving current directory: %s", err.Error())
		os.Exit(2)
	}
	cmdExec := exec.Command("sh", "-c", "git commit --amend -C HEAD")
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	fmt.Println(currentDirectory)
	cmdExec.Run()
}
