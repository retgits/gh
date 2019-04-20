// Package github contains the methods to connect to GitHub as a version control system
package github

import (
	"context"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

// System contains the fields requires to connect to GitHub
type System struct {
	AccessToken string
}

func newClient(ctx context.Context, gh System) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gh.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

// CreateRepository creates a new repository in GitHub
func (gh System) CreateRepository(name string, organization string, private bool) error {
	ctx := context.Background()
	client := newClient(ctx, gh)

	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(private),
	}

	_, _, err := client.Repositories.Create(ctx, organization, repo)

	return err
}
