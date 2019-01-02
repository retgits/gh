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

// gitUndoCmd represents the undo command
var gitUndoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last commit, but don't throw away any changes",
	Run:   runGitUndo,
}

// Flags
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitUndoCmd)
}

// runGitUndo is the actual execution of the command
func runGitUndo(cmd *cobra.Command, args []string) {
	currentDirectory, err := util.GetCurrentDirectory()
	if err != nil {
		fmt.Printf("An error occurred while resolving current directory: %s", err.Error())
		os.Exit(2)
	}
	cmdExec := exec.Command("sh", "-c", "git reset --soft HEAD^")
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	fmt.Println(currentDirectory)
	cmdExec.Run()
}
