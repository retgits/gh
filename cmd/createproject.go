// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/retgits/gh/templates"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Create a Go project",
	Run:   createProject,
}

var (
	author  string
	project string
	base    string
)

func init() {
	rootCmd.AddCommand(createProjectCmd)
	createProjectCmd.Flags().StringVar(&author, "author", "", "The author for the project (required, overrides git.author from config)")
	createProjectCmd.Flags().StringVar(&project, "project", "", "The name of the project (defaults to the current directory)")
	createProjectCmd.Flags().StringVar(&base, "base", "", "The base path of the project (defaults to the current directory)")
	viper.BindPFlag("git.author", createProjectCmd.Flags().Lookup("author"))
}

func createProject(cmd *cobra.Command, args []string) {
	author = viper.GetString("git.author")
	if len(author) == 0 {
		fmt.Printf("author not set as flag (--author) or config (git.author)")
		return
	}

	// Get the current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	if len(project) == 0 {
		project = filepath.Base(dir)
	}

	if len(base) == 0 {
		base = dir
	}

	data := make(map[string]string)
	data["Author"] = author
	data["Project"] = project
	data["Date"] = time.Now().Format("2006-01-02")

	err = executeTemplate("dockerfile", "Dockerfile", templates.Dockerfile, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = executeTemplate("makefile", "Makefile", templates.Makefile, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = executeTemplate("main.go", "main.go", templates.Main, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func executeTemplate(typeName string, fileName string, tmpl string, data map[string]string) error {
	rt, err := template.New(typeName).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error reading %s template: %s", typeName, err.Error())
	}

	rm, err := os.Create(filepath.Join(base, fileName))
	if err != nil {
		return fmt.Errorf("error creating %s: %s", typeName, err.Error())
	}
	defer rm.Close()

	err = rt.Execute(rm, data)
	if err != nil {
		return fmt.Errorf("error executing %s template: %s", typeName, err.Error())
	}

	return nil
}
