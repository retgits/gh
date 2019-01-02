// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Create a GitHub repository",
	Run:   runGitHub,
}

// Flags
var (
	ghtoken string
	ghrepo  string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.Flags().StringVar(&ghtoken, "ghtoken", "", "The Personal Access Token for GitHub (this flag overrides git.ghtoken from the configuration file)")
	githubCmd.Flags().StringVar(&ghrepo, "ghrepo", "", "The repository name to create (will default to the name of the directory if not set)")
	viper.BindPFlag("git.ghtoken", cloneCmd.Flags().Lookup("ghtoken"))
}

// runGitHub is the actual execution of the command
func runGitHub(cmd *cobra.Command, args []string) {
	// Set the github token to use
	ghtoken = viper.GetString("git.ghtoken")
	if len(ghtoken) == 0 {
		fmt.Printf("no ghtoken set in .ghconfig and no --ghtoken flag specified\n%s", cmd.Long)
		os.Exit(1)
	}

	// Set the name of the repo to create
	if len(ghrepo) == 0 {
		// Get the current directory
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		// Get the base directory
		ghrepo = filepath.Base(dir)
	}

	createRepository(ghrepo, "https://api.github.com/user/repos", ghtoken, "github")
}
