// Package gogs contains the methods to connect to Gogs as a version control system
package gogs

import (
	"fmt"
	"net/http"
	"strings"
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

	resp, err := post(url, "application/json", payload, httpHeader)
	if err != nil {
		return fmt.Errorf("error: Gogs did not respond with HTTP/201\n  HTTP StatusCode %v\n  HTTP Body %v", resp, err.Error())
	}

	return nil
}

func post(URL string, encoding string, postData string, header http.Header) (int, error) {
	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(postData))
	if err != nil {
		return 0, fmt.Errorf("error while creating HTTP request: %s", err.Error())
	}

	// Associate the headers with the request
	if header != nil {
		req.Header = header
	}

	// Execute the HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode < 200 || res.StatusCode > 299 {
		return 0, fmt.Errorf("error while performing HTTP request: %s", err.Error())
	}

	return res.StatusCode, nil
}
