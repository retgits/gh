// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/retgits/gh/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The constants
const (
	jenkinsDSLTemplate = `// Project information
	String project = "{{.reponame}}"
	String icon = "search.png"
	
	// GitHub information
	String gitHubRepository = "{{.reponame}}"
	String gitHubUser = "retgits"
	
	// Gogs information
	String gogsRepository = "{{.reponame}}"
	String gogsUser = "retgits"
	String gogsHost = "ubusrvls.na.tibco.com:3000"
	
	// Job DSL definition
	freeStyleJob("mirror-$project") {
	 displayName("mirror-$project")
	 description("Mirror github.com/$gitHubUser/$gitHubRepository")
	
	 checkoutRetryCount(3)
	
	 properties {
	  githubProjectUrl("https://github.com/$gitHubUser/$gitHubRepository")
	  sidebarLinks {
	   link("http://$gogsHost/$gogsUser/$gogsRepository", "Gogs", "$icon")
	  }
	 }
	
	 logRotator {
	  numToKeep(100)
	  daysToKeep(15)
	 }
	
	 triggers {
	  cron('@daily')
	 }
	
	 wrappers {
	  colorizeOutput()
	  credentialsBinding {
	   usernamePassword('GOGS_USERPASS', 'gogs')
	  }
	 }
	
	 steps {
	  shell("git clone --mirror https://github.com/$gitHubUser/$gitHubRepository repo")
	  shell("cd repo && git push --mirror http://\$GOGS_USERPASS@gogs:3000/$gogsUser/$gogsRepository")
	 }
	
	 publishers {
	  mailer {
	   recipients('$ADMIN_EMAIL')
	   notifyEveryUnstableBuild(true)
	   sendToIndividuals(false)
	  }
	  wsCleanup()
	 }
	}`
)

// jenkinsCmd represents the jenkins command
var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Create a Jenkins Job",
	Run:   runGit,
}

// Flags
var (
	commit      bool
	jenkinsrepo string
	projectname string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(jenkinsCmd)
	jenkinsCmd.Flags().BoolVar(&commit, "commit", false, "Commit and push the updates to the Jenkins DSL project")
	jenkinsCmd.Flags().StringVar(&jenkinsrepo, "jenkinsrepo", "", "The location of the Jenkins Job DSL project (this flag overrides git.jenkinsrepo from the configuration file)")
	jenkinsCmd.Flags().StringVar(&projectname, "projectname", "", "The name of the project to create a new job for (will default to the name of the directory if not set)")
	viper.BindPFlag("git.jenkinsrepo", cloneCmd.Flags().Lookup("jenkinsrepo"))
}

// runGit is the actual execution of the command
func runGit(cmd *cobra.Command, args []string) {
	// Set the basefolder to clone to
	jenkinsrepo = viper.GetString("git.jenkinsrepo")
	if len(jenkinsrepo) == 0 {
		fmt.Printf("no jenkinsrepo set in .ghconfig and no --jenkinsrepo flag specified\n%s", cmd.Long)
		os.Exit(1)
	}

	// Set the name of the project to create
	if len(projectname) == 0 {
		// Get the current directory
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		// Get the base directory
		projectname = filepath.Base(dir)
	}

	// Prepare a map
	data := make(map[string]interface{})
	data["reponame"] = projectname

	// Prepare the template
	t := template.Must(template.New("email").Parse(jenkinsDSLTemplate))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		fmt.Printf("there was a problem creating the Jenkins template\n%s\n", err.Error())
	}
	s := buf.String()

	// Write the template to disk
	err := util.WriteFile(filepath.Join(jenkinsrepo, "projects", fmt.Sprintf("%s.groovy", projectname)), s)
	if err != nil {
		fmt.Printf("there was a problem syncing the template file\n%s\n", err.Error())
	}

	// Push to master repo
	if commit {
		cmd := exec.Command("git", "add", "-A", ".", "&&", "git", "commit", "-a", "-m", fmt.Sprintf("\"Add new job for %s\"", projectname), "&&", "git", "push", "origin", "master")
		cmd.Dir = jenkinsrepo
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("rhere was a problem pushing to master branch\n%s\n", err.Error())
		}
	}
}
