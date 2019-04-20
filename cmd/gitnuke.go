// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
)

var gitNukeCmd = &cobra.Command{
	Use:   "nuke-branch",
	Short: "Removes a branch locally and on the remote origin",
	Run:   gitNuke,
}

var (
	gitNukeBranch string
)

func init() {
	rootCmd.AddCommand(gitNukeCmd)
	gitNukeCmd.Flags().StringVar(&gitNukeBranch, "branch", "", "The branch to remove (required)")
	gitNukeCmd.MarkFlagRequired("branch")
}

func gitNuke(cmd *cobra.Command, args []string) {
	err := exec.RunCmd(fmt.Sprintf("git branch -D %s", gitNukeBranch))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = exec.RunCmd(fmt.Sprintf("git push origin :%s", gitNukeBranch))
	if err != nil {
		fmt.Println(err.Error())
	}
}
