// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/retgits/gh/util"
	"github.com/spf13/cobra"
)

// travisCmd represents the travis command
var travisCmd = &cobra.Command{
	Use:   "travis",
	Short: "a command to update the AWS credentials on Travis-CI jobs.",
	Run:   runTravis,
}

// Flags
var (
	travisToken string
	repoList    string
	repoOwner   string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(travisCmd)
	travisCmd.Flags().StringVar(&travisToken, "travis-token", "", "The Authentication Token for Travis-CI (optional)")
	travisCmd.Flags().StringVar(&repoList, "repos", "", "The list of Travis-CI repos to update (optional, must be a comma separated list)")
	travisCmd.Flags().StringVar(&repoOwner, "owner", "", "The owner of the Travis-CI repos (required)")
	travisCmd.MarkFlagRequired("owner")
}

// runTravis is the actual execution of the command
func runTravis(cmd *cobra.Command, args []string) {
	// Get the Travis token. The precedence is as follows:
	// 1) Flag   : travis-token
	// 2) Env var: TRAVISTOKEN
	if len(travisToken) == 0 {
		travisToken = os.Getenv("TRAVISTOKEN")
		if len(travisToken) == 0 {
			fmt.Println("Cannot find Travis-CI token from flags or environment")
		}
	}

	// Get the new value for AWS_ACCESS_KEY_ID
	command := exec.Command("aws", "configure", "get", "aws_access_key_id")
	output, err := command.Output()
	if err != nil {
		fmt.Printf("There was a problem getting the AWS_ACCESS_KEY_ID\n%s\n", err.Error())
	}
	accessKey := string(output)

	// Get the new value for AWS_SECRET_ACCESS_KEY
	command = exec.Command("aws", "configure", "get", "aws_secret_access_key")
	output, err = command.Output()
	if err != nil {
		fmt.Printf("There was a problem getting the AWS_SECRET_ACCESS_KEY\n%s\n", err.Error())
	}
	secretKey := string(output)

	// Get the repo list. The precedence is as follows:
	// 1) Flag   : repos
	// 2) Env var: TRAVISREPOS
	if len(repoList) == 0 {
		repoList = os.Getenv("TRAVISREPOS")
		if len(repoList) == 0 {
			fmt.Println("Cannot find a list of TravisCI repositories to update")
		}
	}

	// Split the repos
	repos := strings.Split(repoList, ",")

	// Get the environment variables for each repository
	for _, repo := range repos {
		// Prepare the HTTP headers
		httpHeader := http.Header{"Authorization": {fmt.Sprintf("token %s", travisToken)}, "Travis-API-Version": {"3"}}

		// Send the API call
		resp, err := util.HTTPGet(fmt.Sprintf("https://api.travis-ci.org/repo/%s%s%s/env_vars", repoOwner, "%2F", repo), "application/json", httpHeader)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Create an array of env vars
		envArray := resp.Body["env_vars"].([]interface{})

		// Loop over the env vars and see if they need to be updated
		for _, envVar := range envArray {
			envVar := envVar.(map[string]interface{})
			if envVar["name"] == "AWS_ACCESS_KEY_ID" {
				// Prepare the payload
				jsonString := fmt.Sprintf(`{"env_var.value":"%s","env_var.public":"false"}`, accessKey)
				// Prepare the HTTP headers
				httpHeader := http.Header{"Authorization": {fmt.Sprintf("token %s", travisToken)}, "Travis-API-Version": {"3"}}
				// Send the API call
				resp, err := util.HTTPPatch(fmt.Sprintf("https://api.travis-ci.org/repo/%s%s%s/env_var/%s", repoOwner, "%2F", repo, envVar["id"].(string)), "application/json", jsonString, httpHeader)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("Update of AWS_ACCESS_KEY_ID resulted in statuscode %d\n", resp.StatusCode)
			} else if envVar["name"] == "AWS_SECRET_ACCESS_KEY" {
				// Prepare the payload
				jsonString := fmt.Sprintf(`{"env_var.value":"%s","env_var.public":"false"}`, secretKey)
				// Prepare the HTTP headers
				httpHeader := http.Header{"Authorization": {fmt.Sprintf("token %s", travisToken)}, "Travis-API-Version": {"3"}}
				// Send the API call
				resp, err := util.HTTPPatch(fmt.Sprintf("https://api.travis-ci.org/repo/%s%s%s/env_var/%s", repoOwner, "%2F", repo, envVar["id"].(string)), "application/json", jsonString, httpHeader)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("Update of AWS_SECRET_ACCESS_KEY resulted in statuscode %d\n", resp.StatusCode)
			}
		}
	}
}
