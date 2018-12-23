// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/util"
	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Run git commands",
	Run:   runGit,
}

// Flags
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitCmd)
}

// runGit is the actual execution of the command
func runGit(cmd *cobra.Command, args []string) {
	fmt.Printf("\nThe git allows for certain subcommands.\nThe available subcommands are:\n\n")

	// Print all subcommands
	for _, command := range cmd.Commands() {
		if command.Use != "help [command]" {
			fmt.Printf("%s %s\n", util.RightPadToLen(command.Use, ".", 15), command.Short)
		}
	}

	fmt.Println("")
}
