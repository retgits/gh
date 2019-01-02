// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/retgits/gh/util"

	"github.com/spf13/cobra"
)

// gitNukeCmd represents the nuke command
var gitNukeCmd = &cobra.Command{
	Use:   "nuke",
	Short: "Removes a branch locally and on the remote origin",
	Run:   runGitNuke,
}

// Flags
var (
	gitNukeBranch string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitNukeCmd)
	gitNukeCmd.Flags().StringVar(&gitNukeBranch, "branch", "", "The Nuke message (required)")
	gitNukeCmd.MarkFlagRequired("branch")
}

// runGitNuke is the actual execution of the command
func runGitNuke(cmd *cobra.Command, args []string) {
	currentDirectory, err := util.GetCurrentDirectory()
	if err != nil {
		fmt.Printf("An error occurred while resolving current directory: %s", err.Error())
		os.Exit(2)
	}
	cmdExec := exec.Command("sh", "-c", fmt.Sprintf("git branch -D %s", gitNukeBranch))
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	cmdExec.Run()

	cmdExec = exec.Command("sh", "-c", fmt.Sprintf("git push origin :%s", gitNukeBranch))
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	cmdExec.Run()
}
