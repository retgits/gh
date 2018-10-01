// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "Git Helper",
	Long: `
A collection of git helper commands to make my life a little easier`,
}

// Variables used in multiple flags
var (
	base        string
	githubToken string
	repoOwner   string
)

const (
	version = "1.3.0"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("\nYou're running gh version {{.Version}}\n\n")
}
