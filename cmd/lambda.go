// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/retgits/gh/util"
	"github.com/spf13/cobra"
)

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "Create a new AWS Lambda function based on my personal templates in the current folder.",
	Run:   runLambda,
	Long:  "\ngh lambda is a command to create a new AWS Lambda function based on my personal templates in the current folder\n\nSample usage: gh lambda my-lambda\nThis will create a new AWS Lambda function in the my-lambda folder of this directory\n\n",
}

// Flags
var (
	name string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(lambdaCmd)
	lambdaCmd.Flags().StringVar(&name, "name", "", "The name of the lambda function you want to create (required)")
	lambdaCmd.Flags().StringVar(&base, "base", "", "The root folder to create this lambda function in (optional, will default to current folder)")
	lambdaCmd.MarkFlagRequired("name")
}

// runLambda is the actual execution of the command
func runLambda(cmd *cobra.Command, args []string) {
	// Set base to the current folder if it wasn't specified as a command line argument
	if len(base) == 0 {
		base, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	// Create the function folder
	folder := filepath.Join(base, name)
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// Create the test folder
	folder = filepath.Join(base, name, "test")
	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// Set folder back to it's original setting
	folder = filepath.Join(base, name)

	// Prepare a map with data
	dataMap := make(map[string]interface{})
	dataMap["name"] = name

	// Write the templates
	err = util.WriteFile(filepath.Join(folder, ".gitignore"), gitIgnore)
	if err != nil {
		fmt.Println(err)
	}

	err = util.WriteFile(filepath.Join(folder, "main.go"), mainGo)
	if err != nil {
		fmt.Println(err)
	}

	t := template.Must(template.New("top").Parse(makefile))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, dataMap); err != nil {
		fmt.Printf("error while rendering makefile: %s", err.Error())
	}

	err = util.WriteFile(filepath.Join(folder, "makefile"), buf.String())
	if err != nil {
		fmt.Println(err)
	}

	t = template.Must(template.New("top").Parse(yamlTemplate))
	buf = &bytes.Buffer{}
	if err := t.Execute(buf, dataMap); err != nil {
		fmt.Printf("error while rendering makefile: %s", err.Error())
	}

	err = util.WriteFile(filepath.Join(folder, "template.yml"), buf.String())
	if err != nil {
		fmt.Println(err)
	}

	err = util.WriteFile(filepath.Join(folder, "LICENSE"), license)
	if err != nil {
		fmt.Println(err)
	}
}
