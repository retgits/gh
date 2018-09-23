// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "a simple git clone command to make sure that all git clones end up in a specified directory.",
	Run:   runClone,
	Long:  "\ngh clone is a simple git clone command to make sure that all git clones end up in a specified\ndirectory. The directory is specified by\n1) setting a flag `base` (gh clone --base . https://github.com/retgits/gh)\n2) setting an environment variable `GITBASEFOLDER`\n3) the current directory\n\nSample usage: gh clone https://github.com/retgits/gh\n\n",
}

// Flags
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVar(&base, "base", "", "The root folder to clone this repo in (optional, unless $GITBASEFOLDER is set)")
}

// runClone is the actual execution of the command
func runClone(cmd *cobra.Command, args []string) {
	// If the URL isn't provided as a commandline argument, stop processing
	if len(args) < 1 {
		fmt.Printf("Error: There was no URL provided\n%s", cmd.Long)
		os.Exit(1)
	}

	// If the base flag wasn't specified check if an environment variable was set
	if len(base) == 0 {
		// Get the base directory
		envVar, set := os.LookupEnv("GITBASEFOLDER")
		if !set {
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				fmt.Printf("Error: %s\n%s", err.Error(), cmd.Long)
				os.Exit(1)
			}
			base = dir
		} else {
			base = envVar
		}
	}

	// Get the git URL and split it into URL segments
	gitURL := strings.Split(args[0], "/")
	if len(gitURL) < 4 {
		fmt.Printf("Error: Not enough arguments in %v\n%s", gitURL, cmd.Long)
		os.Exit(1)
	}
	// The gitOrigin represents the domain name of the git server (like github.com)
	gitOrigin := gitURL[2]
	// The gitUser respresents the repository owner (like retgits)
	gitUser := gitURL[3]
	// The gitRepo is the actual name of the repository
	gitRepo := strings.Replace(gitURL[4], ".git", "", -1)

	// localDir is the location on disk that the repository will be cloned to
	localDir := filepath.Join(base, gitOrigin, gitUser, gitRepo)

	// Clone the repository
	command := exec.Command("git", "clone", args[0], localDir)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Printf("Error: %s\n%s", err.Error(), cmd.Long)
		os.Exit(1)
	}
}
