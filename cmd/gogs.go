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

// gogsCmd represents the gogs command
var gogsCmd = &cobra.Command{
	Use:   "gogs",
	Short: "Create a Gogs repository",
	Run:   runGogs,
}

// Flags
var (
	gogstoken string
	gogsrepo  string
	gogsurl   string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gogsCmd)
	gogsCmd.Flags().StringVar(&gogstoken, "gogstoken", "", "The Personal Access Token for gogs (this flag overrides git.gogstoken from the configuration file)")
	gogsCmd.Flags().StringVar(&gogsrepo, "gogsrepo", "", "The repository name to create (will default to the name of the directory if not set)")
	gogsCmd.Flags().StringVar(&gogsurl, "gogsurl", "", "The URL of the gogs server (this flag overrides git.gogsurl from the configuration file)")
	viper.BindPFlag("git.gogstoken", cloneCmd.Flags().Lookup("gogstoken"))
	viper.BindPFlag("git.gogsurl", cloneCmd.Flags().Lookup("gogsurl"))
}

// runGogs is the actual execution of the command
func runGogs(cmd *cobra.Command, args []string) {
	// Set the gogs token to use
	gogstoken = viper.GetString("git.gogstoken")
	if len(gogstoken) == 0 {
		fmt.Printf("no ghtoken set in .ghconfig and no --gogstoken flag specified\n%s", cmd.Long)
		os.Exit(1)
	}

	// Set the gogs url to use
	gogsurl = viper.GetString("git.gogsurl")
	if len(gogstoken) == 0 {
		fmt.Printf("no gogsurl set in .ghconfig and no --gogsurl flag specified\n%s", cmd.Long)
		os.Exit(1)
	}

	// Set the name of the repo to create
	if len(gogsrepo) == 0 {
		// Get the current directory
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		// Get the base directory
		gogsrepo = filepath.Base(dir)
	}

	createRepository(gogsrepo, gogsurl, gogstoken, "gogs")
}
