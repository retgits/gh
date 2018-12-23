// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"net/http"

	"github.com/retgits/gh/util"
)

func createRepository(reponame string, URL string, token string, origin string) {
	// Prepare the payload
	jsonString := fmt.Sprintf(`{"name":"%s"}`, reponame)

	// Prepare the HTTP headers
	httpHeader := http.Header{"Authorization": {fmt.Sprintf("token %s", token)}}

	// Send the API call
	resp, err := util.HTTPPost(URL, "application/json", jsonString, httpHeader)
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode != 201 {
		fmt.Printf("%s did not respond with HTTP/201\n", origin)
		fmt.Printf("  HTTP StatusCode %v\n", resp.StatusCode)
		fmt.Printf("  HTTP Body %v\n", resp.Body)
	}

	fmt.Println(resp.Body)
}
