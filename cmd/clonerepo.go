// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository to a specified directory",
	Run:   runClone,
	Long:  "clone makes sure repositories are cloned to a specified base directory and a predefined structure:\n  <basefolder>/<git site>/<user>/<repo> (like /home/user/github.com/retgits/gh). The basefolder is\n  set by the git.basefolder in .ghconfig.yml or the --basefolder flag\n\nsample usage: gh clone https://github.com/retgits/gh\n\n",
}

// Flags
var (
	basefolder string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringVar(&basefolder, "basefolder", "", "The root folder to clone to")
	viper.BindPFlag("git.basefolder", cloneCmd.Flags().Lookup("basefolder"))
}

// runClone is the actual execution of the command
func runClone(cmd *cobra.Command, args []string) {
	// If the URL isn't provided as a commandline argument, stop processing
	if len(args) < 1 {
		fmt.Printf("no URL provided to clone\n")
		os.Exit(1)
	}

	// Get the git URL and split it into URL segments
	gitURL := strings.Split(args[0], "/")
	if len(gitURL) < 4 {
		fmt.Printf("not enough arguments provided in %v\n", gitURL)
		os.Exit(1)
	}

	// Set the basefolder to clone to
	basefolder = viper.GetString("git.basefolder")
	if len(basefolder) == 0 {
		basefolder = "."
	}

	// localDir is the location on disk that the repository will be cloned to
	// gitURL[2] is the domain name of the git server (like github.com)
	// gitURL[3] is the repository owner (like retgits)
	// gitURL[4] is the actual name of the repository
	localDir := filepath.Join(basefolder, gitURL[2], gitURL[3], strings.Replace(gitURL[4], ".git", "", -1))

	// Clone the repository
	err := exec.RunCmd(fmt.Sprintf("git clone %s %s", args[0], localDir))
	if err != nil {
		fmt.Printf("error running git clone command: %s\n", err.Error())
	}
}
