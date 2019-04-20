// Package gogs contains the methods to connect to Gogs as a version control system
package gogs

import (
	"fmt"
	"net/http"

	"github.com/retgits/gh/util"
)

// System contains the fields requires to connect to Gogs
type System struct {
	AccessToken string
	Endpoint    string
}

// CreateRepository creates a new repository in Gogs
func (g System) CreateRepository(name string, organization string, private bool) error {
	url := fmt.Sprintf("%s/user/repos", g.Endpoint)
	if len(organization) != 0 {
		url = fmt.Sprintf("%s/orgs/%s/repos", g.Endpoint, organization)
	}

	fmt.Println(url)

	payload := fmt.Sprintf(`{"name":"%s","private":"%v"}`, name, private)

	httpHeader := http.Header{"Authorization": {fmt.Sprintf("token %s", g.AccessToken)}}

	resp, err := util.HTTPPost(url, "application/json", payload, httpHeader)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Gogs did not respond with HTTP/201\n  HTTP StatusCode %v\n  HTTP Body %v", resp.StatusCode, resp.Body)
	}

	return nil
}
