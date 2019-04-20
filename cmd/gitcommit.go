// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
)

// gitCommitCmd represents the commit command
var gitCommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A simpler alias for \"git commit -a -S -m\"",
	Run:   gitCommit,
}

var (
	gitCommitMessage string
)

func init() {
	rootCmd.AddCommand(gitCommitCmd)
	gitCommitCmd.Flags().StringVar(&gitCommitMessage, "message", "", "The commit message (required)")
	gitCommitCmd.MarkFlagRequired("message")
}

func gitCommit(cmd *cobra.Command, args []string) {
	err := exec.RunCmd(fmt.Sprintf("git commit -a -S -m \"%s\"", gitCommitMessage))
	if err != nil {
		fmt.Println(err.Error())
	}
}
