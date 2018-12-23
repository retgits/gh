// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/retgits/gh/util"
	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Create GitHub and/or Gogs repositories and optionally a Jenkins job as well.",
	Run:   runGit,
}

// Flags
var (
	ghub        bool
	gogs        bool
	jenkins     bool
	commit      bool
	gogsToken   string
	jenkinsBase string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitCmd)
	gitCmd.Flags().BoolVar(&ghub, "github", false, "Create a GitHub repository for this project")
	gitCmd.Flags().BoolVar(&gogs, "gogs", false, "Create a Gogs repository for this project")
	gitCmd.Flags().BoolVar(&jenkins, "jenkins", false, "Create a Jenkins DSL for this project")
	gitCmd.Flags().BoolVar(&commit, "commit", false, "Commit and push the updates to the Jenkins DSL project")
	gitCmd.Flags().StringVar(&githubToken, "github-token", "", "The Personal Access Token for GitHub (optional)")
	gitCmd.Flags().StringVar(&gogsToken, "gogs-token", "", "The Personal Access Token for Gogs (optional)")
	gitCmd.Flags().StringVar(&jenkinsBase, "jenkins-base", "", "The base directory of the Jenkins DSL project (optional)")
}

// runGit is the actual execution of the command
func runGit(cmd *cobra.Command, args []string) {
	// Get the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	// Get the base directory
	base := filepath.Base(dir)

	// Get the Gogs token. The precedence is as follows:
	// 1) Flag   : gogs-token
	// 2) Env var: GOGSTOKEN
	if gogs {
		if len(gogsToken) == 0 {
			gogsToken = os.Getenv("GOGSTOKEN")
			if len(gogsToken) == 0 {
				fmt.Println("Cannot find Gogs token from flags or environment")
			}
		}
		createRepository(base, gogsToken, "gogs")
	}

	// Get the GitHub token. The precedence is as follows:
	// 1) Flag   : github-token
	// 2) Env var: GITHUBTOKEN
	if ghub {
		if len(githubToken) == 0 {
			githubToken = os.Getenv("GITHUBTOKEN")
			if len(githubToken) == 0 {
				fmt.Println("Cannot find GitHub token from flags or environment")
			}
		}
		createRepository(base, githubToken, "github")
	}

	// Create a Jenkins job.The precedence is as follows:
	// 1) Flag   : jenkins-base
	// 2) Env var: JENKINSBASEDIR
	if jenkins {
		if len(jenkinsBase) == 0 {
			jenkinsBase = os.Getenv("JENKINSBASEDIR")
			if len(jenkinsBase) == 0 {
				fmt.Println("Cannot find Jenkins base directory from flags or environment")
			}
		}
		createJenkinsJob(base, jenkinsBase, commit)
	}
}

func createJenkinsJob(reponame string, jenkinsBase string, commit bool) {
	// Prepare a map
	data := make(map[string]interface{})
	data["reponame"] = reponame

	// Prepare the template
	t := template.Must(template.New("email").Parse(jenkinsDSLTemplate))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		fmt.Printf("There was a problem creating the Jenkins template\n%s\n", err.Error())
	}
	s := buf.String()

	// Write the template to disk
	err := util.WriteFile(filepath.Join(jenkinsBase, "projects", fmt.Sprintf("%s.groovy", reponame)), s)
	if err != nil {
		fmt.Printf("There was a problem syncing the template file\n%s\n", err.Error())
	}

	// Push to GitHub
	if commit {
		cmd := exec.Command("git", "add", "-A", ".", "&&", "git", "commit", "-a", "-m", fmt.Sprintf("\"Add new job for %s\"", reponame), "&&", "git", "push", "origin", "master")
		cmd.Dir = jenkinsBase
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("There was a problem pushing to GitHub\n%s\n", err.Error())
		}
	}
}

func createRepository(reponame string, token string, origin string) {
	// Prepare the payload
	jsonString := fmt.Sprintf(`{"name":"%s"}`, reponame)

	// Prepare the HTTP headers
	httpHeader := http.Header{"Authorization": {fmt.Sprintf("token %s", token)}}

	// Set the URL based on the origin
	var URL string
	if origin == "github" {
		URL = "https://api.github.com/user/repos"
	} else {
		URL = "http://ubusrvls.na.tibco.com:3000/api/v1/user/repos"
	}

	// Send the API call
	resp, err := util.HTTPPost(URL, "application/json", jsonString, httpHeader)
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode != 201 {
		fmt.Printf("GitHub did not response with HTTP/201\n")
		fmt.Printf("  HTTP StatusCode %v\n", resp.StatusCode)
		fmt.Printf("  HTTP Body %v\n", resp.Body)
	}

	fmt.Println(resp.Body)
}
