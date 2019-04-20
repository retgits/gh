// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/gh/exec"
	"github.com/spf13/cobra"
)

var gitCreditCmd = &cobra.Command{
	Use:   "credit",
	Short: "A very slightly quicker way to credit an author on the latest commit",
	Run:   gitCredit,
}

var (
	gitCreditName  string
	gitCreditEmail string
)

func init() {
	rootCmd.AddCommand(gitCreditCmd)
	gitCreditCmd.Flags().StringVar(&gitCreditName, "name", "", "The name of the author to credit (required)")
	gitCreditCmd.Flags().StringVar(&gitCreditEmail, "email", "", "The email address of the author to credit (required)")
	gitCreditCmd.MarkFlagRequired("name")
	gitCreditCmd.MarkFlagRequired("email")
}

func gitCredit(cmd *cobra.Command, args []string) {
	err := exec.RunCmd(fmt.Sprintf("git commit --amend --author \"%s <%s>\" -C HEAD", gitCreditName, gitCreditEmail))
	if err != nil {
		fmt.Println(err.Error())
	}
}
