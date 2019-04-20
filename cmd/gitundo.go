// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
)

var gitUndoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last commit, but don't throw away any changes",
	Run:   gitUndo,
}

func init() {
	rootCmd.AddCommand(gitUndoCmd)
}

func gitUndo(cmd *cobra.Command, args []string) {
	err := exec.RunCmd("git reset --soft HEAD^")
	if err != nil {
		fmt.Println(err.Error())
	}
}
