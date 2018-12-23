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

// gitCommitCmd represents the git command
var gitCommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A simpler alias for \"git commit -a -S -m\"",
	Run:   runGitCommit,
}

// Flags
var (
	gitCommitMessage string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitCommitCmd)
	gitCommitCmd.Flags().StringVar(&gitCommitMessage, "message", "", "The commit message (required)")
	gitCommitCmd.MarkFlagRequired("message")
}

// runGitCommit is the actual execution of the command
func runGitCommit(cmd *cobra.Command, args []string) {
	currentDirectory, err := util.GetCurrentDirectory()
	if err != nil {
		fmt.Printf("An error occured while resolving current directory: %s", err.Error())
		os.Exit(2)
	}
	cmdExec := exec.Command("sh", "-c", fmt.Sprintf("git commit -a -S -m \"%s\"", gitCommitMessage))
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	fmt.Println(currentDirectory)
	cmdExec.Run()
}
