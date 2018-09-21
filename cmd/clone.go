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
var (
	base string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVar(&base, "base", "", "The root folder to clone this repo in (optional, unless $GITBASEFOLDER is set)")
}

// runClone is the actual execution of the command
func runClone(cmd *cobra.Command, args []string) {

	// If the base flag wasn't specified check if an environment variable was set
	if len(base) == 0 {
		// Get the base directory
		envVar, set := os.LookupEnv("GITBASEFOLDER")
		if !set {
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				fmt.Println(cmd.Long)
				os.Exit(1)
			}
			base = dir
		} else {
			base = envVar
		}
	}

	// Get the git URL
	gitURL := strings.Split(os.Args[len(os.Args)-1], "/")
	if len(gitURL) < 4 {
		fmt.Printf("Not enough arguments in %v\n\n", gitURL)
		fmt.Println(cmd.Long)
		os.Exit(1)
	}
	gitOrigin := gitURL[2]
	gitUser := gitURL[3]
	gitRepo := strings.Replace(gitURL[4], ".git", "", -1)

	// Prepare the location
	localDir := filepath.Join(base, gitOrigin, gitUser, gitRepo)

	// Clone the repository
	fmt.Printf("git clone %s %s\n", os.Args[len(os.Args)-1], localDir)
	command := exec.Command("git", "clone", os.Args[len(os.Args)-1], localDir)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Println(cmd.Long)
		os.Exit(1)
	}
}
