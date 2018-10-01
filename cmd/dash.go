// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v18/github"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	// The database is sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// dashCmd represents the dash command
var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "a command to update the snippets in Dash with GitHub gists.",
	Run:   runDash,
}

// Flags
var (
	dashLib string
)

// Constants
const (
	gistsURL = "https://api.github.com/gists/"
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(dashCmd)
	dashCmd.Flags().StringVar(&githubToken, "github-token", "", "The Personal Access Token for GitHub (optional)")
	dashCmd.Flags().StringVar(&dashLib, "lib", "", "The full path to the library.dash file (like /Users/username/Library/Application Support/Dash/library.dash)")
	dashCmd.Flags().StringVar(&repoOwner, "owner", "", "The GitHub username to get gists for (required)")
	dashCmd.MarkFlagRequired("owner")
	dashCmd.MarkFlagRequired("lib")
}

// runDash is the actual execution of the command
func runDash(cmd *cobra.Command, args []string) {
	// Get the GitHub token. The precedence is as follows:
	// 1) Flag   : github-token
	// 2) Env var: GITHUBTOKEN
	if len(githubToken) == 0 {
		githubToken = os.Getenv("GITHUBTOKEN")
		if len(githubToken) == 0 {
			fmt.Println("Cannot find GitHub token from flags or environment")
		}
	}

	// Open a connection to the database
	dbase, err := sqlx.Open("sqlite3", dashLib)
	if err != nil {
		fmt.Printf("Error while opening Dash library: %s\n", err.Error())
		os.Exit(1)
	}
	defer dbase.Close()

	// Get all current snippets from the Dash library
	dashSnippets := make(map[string]string)
	rows, err := dbase.Queryx("SELECT title, sid FROM snippets")
	for rows.Next() {
		var title string
		var sid string
		rows.Scan(&title, &sid)
		dashSnippets[title] = sid
	}

	// Login to GitHub using a Personal Access Token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Get all gists
	gists, _, err := client.Gists.List(ctx, repoOwner, nil)
	if err != nil {
		fmt.Printf("Error while getting gists from GitHub: %s\n", err.Error())
		os.Exit(1)
	}

	// For each gist
	for _, gist := range gists {
		// Only get the data if the gist is not flagged with [skip-dash]
		if !strings.Contains(gist.GetDescription(), "[skip-dash]") {
			// Get the details of the gist
			currentGist, _, err := client.Gists.Get(ctx, gist.GetID())
			if err != nil {
				fmt.Printf("Error while getting gist from GitHub: %s\n", err.Error())
			}

			// Get the content of the gist
			var body string
			for _, file := range currentGist.Files {
				body = fmt.Sprintf("%s\n%s", body, strings.Replace(file.GetContent(), "'", "''", -1))
			}

			// Get all the tags by splitting the description of the gist by a hashtag
			gistTags := strings.Split(currentGist.GetDescription(), "#")

			// If the snippet already exists, do an update else create a new entry
			if sid, ok := dashSnippets[gistTags[0]]; ok {
				_, err := dbase.Exec(fmt.Sprintf("UPDATE snippets SET body = '%s' WHERE sid = '%s';", body, sid))
				if err != nil {
					fmt.Printf("Error while updating gist [%s] in Dash: %s\n", gistTags[0], err.Error())
				}
			} else {
				_, err := dbase.Exec(fmt.Sprintf("INSERT INTO snippets ('sid', 'title', 'body', 'syntax', 'usageCount') VALUES('%d','%s','%s','None','0')", len(dashSnippets), gistTags[0], body))
				if err != nil {
					fmt.Printf("Error while inserting gist [%s] into Dash: %s\n", gistTags[0], err.Error())
				}
				dashSnippets[gistTags[0]] = fmt.Sprintf("%d", len(dashSnippets))
			}
		}
	}
}
