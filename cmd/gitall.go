// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
)

var gitAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Stage all unstaged files",
	Run:   gitAll,
}

func init() {
	rootCmd.AddCommand(gitAllCmd)
}

func gitAll(cmd *cobra.Command, args []string) {
	err := exec.RunCmd("git add -A")
	if err != nil {
		fmt.Println(err.Error())
	}
}
