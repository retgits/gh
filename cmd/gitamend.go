// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
)

var gitAmendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Use the last commit message and amend your stuffs",
	Run:   gitAmend,
}

func init() {
	rootCmd.AddCommand(gitAmendCmd)
}

func gitAmend(cmd *cobra.Command, args []string) {
	err := exec.RunCmd("git commit --amend -C HEAD")
	if err != nil {
		fmt.Println(err.Error())
	}
}
