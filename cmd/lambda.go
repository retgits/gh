// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "a command to create a new AWS Lambda function based on my personal templates in the current folder.",
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
	lambdaCmd.MarkFlagRequired("name")
}

// runLambda is the actual execution of the command
func runLambda(cmd *cobra.Command, args []string) {
	// If the name flag wasn't set
	if len(name) == 0 {
		fmt.Printf("Not enough arguments\n\n")
		fmt.Println(cmd.Long)
		os.Exit(1)
	} else {
		fmt.Printf("Creating new function %s\n\n", name)
	}

	// Get the current folder
	base, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}

	// Create the function folder
	folder := filepath.Join(base, name)
	err = os.MkdirAll(folder, os.ModePerm)
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
	err = writeFile(filepath.Join(folder, ".gitignore"), gitIgnore)
	if err != nil {
		fmt.Println(err)
	}

	err = writeFile(filepath.Join(folder, "main.go"), mainGo)
	if err != nil {
		fmt.Println(err)
	}

	t := template.Must(template.New("top").Parse(makefile))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, dataMap); err != nil {
		fmt.Printf("error while rendering makefile: %s", err.Error())
	}

	err = writeFile(filepath.Join(folder, "makefile"), buf.String())
	if err != nil {
		fmt.Println(err)
	}

	t = template.Must(template.New("top").Parse(yamlTemplate))
	buf = &bytes.Buffer{}
	if err := t.Execute(buf, dataMap); err != nil {
		fmt.Printf("error while rendering makefile: %s", err.Error())
	}

	err = writeFile(filepath.Join(folder, "template.yml"), buf.String())
	if err != nil {
		fmt.Println(err)
	}

	err = writeFile(filepath.Join(folder, "LICENSE"), license)
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile(filename string, content string) error {
	// Create a file on disk
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error while creating file: %s", err.Error())
		return fmt.Errorf("error while creating file: %s", err.Error())
	}
	defer file.Close()

	// Open the file to write
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("error while opening file: %s", err.Error())
		return fmt.Errorf("error while opening file: %s", err.Error())
	}

	// Write the content to disk
	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Printf("error while writing Markdown to disk: %s", err.Error())
		return fmt.Errorf("error while writing Markdown to disk: %s", err.Error())
	}

	return nil
}
