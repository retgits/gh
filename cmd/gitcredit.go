// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/retgits/gh/util"

	"github.com/spf13/cobra"
)

// gitCreditCmd represents the credit command
var gitCreditCmd = &cobra.Command{
	Use:   "credit",
	Short: "A very slightly quicker way to credit an author on the latest commit",
	Run:   runGitCredit,
}

// Flags
var (
	gitCreditName  string
	gitCreditEmail string
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(gitCreditCmd)
	gitCreditCmd.Flags().StringVar(&gitCreditName, "name", "", "The name of the author to credit (required)")
	gitCreditCmd.Flags().StringVar(&gitCreditEmail, "email", "", "The email address of the author to credit (required)")
	gitCreditCmd.MarkFlagRequired("name")
	gitCreditCmd.MarkFlagRequired("email")
}

// runGitCredit is the actual execution of the command
func runGitCredit(cmd *cobra.Command, args []string) {
	currentDirectory, err := util.GetCurrentDirectory()
	if err != nil {
		fmt.Printf("An error occured while resolving current directory: %s", err.Error())
		os.Exit(2)
	}
	cmdExec := exec.Command("sh", "-c", fmt.Sprintf("git commit --amend --author \"%s <%s>\" -C HEAD", gitCreditName, gitCreditEmail))
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = currentDirectory
	fmt.Println(currentDirectory)
	cmdExec.Run()
}
