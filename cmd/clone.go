// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository to a specified directory",
	Run:   runClone,
	Long:  "\nclone makes sure repositories are cloned to a specified base directory and a predefined structure:\n  <basefolder>/<git site>/<user>/<repo> (like /home/user/github.com/retgits/gh). The basefolder is\n  set by the git.basefolder in .ghconfig.yml or the --basefolder flag\n\nsample usage: gh clone https://github.com/retgits/gh\n\n",
}

// Flags
var (
	basefolder string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVar(&basefolder, "basefolder", "", "The root folder to clone to (this flag overrides git.basefolder from the configuration file)")
	viper.BindPFlag("git.basefolder", cloneCmd.Flags().Lookup("basefolder"))
}

// runClone is the actual execution of the command
func runClone(cmd *cobra.Command, args []string) {
	// Set the basefolder to clone to
	basefolder = viper.GetString("git.basefolder")
	if len(basefolder) == 0 {
		fmt.Printf("no basefolder set in .ghconfig and no --basefolder flag specified\n%s", cmd.Long)
		os.Exit(1)
	}

	// If the URL isn't provided as a commandline argument, stop processing
	if len(args) < 1 {
		fmt.Printf("no URL provided to clone\n%s", cmd.Long)
		os.Exit(1)
	}

	// Get the git URL and split it into URL segments
	gitURL := strings.Split(args[0], "/")
	if len(gitURL) < 4 {
		fmt.Printf("not enough arguments provided in %v\n%s", gitURL, cmd.Long)
		os.Exit(1)
	}

	// The gitOrigin represents the domain name of the git server (like github.com)
	gitOrigin := gitURL[2]
	// The gitUser respresents the repository owner (like retgits)
	gitUser := gitURL[3]
	// The gitRepo is the actual name of the repository
	gitRepo := strings.Replace(gitURL[4], ".git", "", -1)

	// localDir is the location on disk that the repository will be cloned to
	localDir := filepath.Join(basefolder, gitOrigin, gitUser, gitRepo)

	// Clone the repository
	command := exec.Command("git", "clone", args[0], localDir)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Printf("error running git clone command: %s\n%s", err.Error(), cmd.Long)
		os.Exit(1)
	}
}
