// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/retgits/gh/vcs"
	"github.com/retgits/gh/vcs/github"
	"github.com/retgits/gh/vcs/gogs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createRepositoryCmd = &cobra.Command{
	Use:   "create-repo",
	Short: "Create a repository",
	Run:   createRepository,
}

var (
	accessToken string
	repoName    string
	orgName     string
	vcsType     string
	vcsURL      string
	privateRepo bool
)

func init() {
	rootCmd.AddCommand(createRepositoryCmd)
	createRepositoryCmd.Flags().StringVar(&accessToken, "token", "", "The Personal Access Token for the version control system")
	createRepositoryCmd.Flags().StringVar(&repoName, "repo", "", "The repository name to create")
	createRepositoryCmd.Flags().StringVar(&orgName, "org", "", "The organization to create the repo under")
	createRepositoryCmd.Flags().StringVar(&vcsType, "type", "", "The version control system to use")
	createRepositoryCmd.Flags().StringVar(&vcsURL, "url", "", "The API endpoint to call")
	createRepositoryCmd.Flags().BoolVar(&privateRepo, "private", false, "Set to true to create a private repository")
}

func createRepository(cmd *cobra.Command, args []string) {
	var client vcs.System

	switch strings.ToLower(vcsType) {
	case "gogs":
		accessToken = getConfigWithFallback(accessToken, "gogs.accesstoken")
		client = gogs.System{AccessToken: accessToken, Endpoint: getConfigWithFallback(vcsURL, "gogs.apiendpoint")}
	case "github":
		accessToken = getConfigWithFallback(accessToken, "github.accesstoken")
		client = github.System{AccessToken: accessToken}
	default:
		fmt.Printf("Unknown version control system %s\n", vcsType)
		os.Exit(1)
	}

	if len(repoName) == 0 {
		// Get the current directory
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		// Get the base directory
		repoName = filepath.Base(dir)
	}

	err := vcs.System.CreateRepository(client, repoName, orgName, privateRepo)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getConfigWithFallback(item string, fallback string) string {
	if len(item) == 0 {
		return viper.GetString(fallback)
	}
	return item
}
