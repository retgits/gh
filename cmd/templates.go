// Package cmd defines and implements command-line commands and flags
// used by fdio. Commands and flags are implemented using Cobra.
package cmd

const (
	gitIgnore = `## Binaries
bin/*
## SAM Local deployment artifacts
packaged.yaml
## Debug
debug/
`

	mainGo = `// Package main is the main implementation of .
package main
// The imports
import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)
// Constants
const (
)
// Variables
var (
)
// The handler function is executed every time that a new Lambda event is received.
// It takes a JSON payload (you can see an example in the event.json file) and only
// returns an error if the something went wrong. The event comes fom CloudWatch and
// is scheduled every interval (where the interval is defined as variable)
func handler(request events.CloudWatchEvent) error {
	return nil
}
// The main method is executed by AWS Lambda and points to the handler
func main() {
	lambda.Start(handler)
}
	`

	makefile = `.PHONY: deps clean build deploy test-lambda
deps:
	go get -u ./...
clean: 
	rm -rf ./bin
	
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/{{.name}} *.go
test-lambda: clean build
	sam local invoke {{.name}} -e ./test/event.json
deploy: clean build
	sam package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket retgits-apps
	sam deploy --template-file packaged.yaml --stack-name {{.name}} --capabilities CAPABILITY_IAM
	`

	yamlTemplate = `
AWSTemplateFormatVersion: '2010-09-09'
Transform: 'AWS::Serverless-2016-10-31'
Description: Serverless Application Model for {{.name}}
Resources:
  {{.name}}:
    Type: 'AWS::Serverless::Function'
    Properties:
      CodeUri: bin/
      Handler: {{.name}}
      Runtime: go1.x
      Tracing: Active
      Timeout: 120
      Tags:
        version: "0.0.1"
      Events:
        {{.name}}:
      Description: App for {{.name}}
      MemorySize: 128`

	license = `MIT License
	Copyright (c) 2018 retgits
	
	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:
	
	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.
	
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
	`

	jenkinsDSLTemplate = `// Project information
	String project = "{{.reponame}}"
	String icon = "search.png"
	
	// GitHub information
	String gitHubRepository = "{{.reponame}}"
	String gitHubUser = "retgits"
	
	// Gogs information
	String gogsRepository = "{{.reponame}}"
	String gogsUser = "retgits"
	String gogsHost = "ubusrvls.na.tibco.com:3000"
	
	// Job DSL definition
	freeStyleJob("mirror-$project") {
	 displayName("mirror-$project")
	 description("Mirror github.com/$gitHubUser/$gitHubRepository")
	
	 checkoutRetryCount(3)
	
	 properties {
	  githubProjectUrl("https://github.com/$gitHubUser/$gitHubRepository")
	  sidebarLinks {
	   link("http://$gogsHost/$gogsUser/$gogsRepository", "Gogs", "$icon")
	  }
	 }
	
	 logRotator {
	  numToKeep(100)
	  daysToKeep(15)
	 }
	
	 triggers {
	  cron('@daily')
	 }
	
	 wrappers {
	  colorizeOutput()
	  credentialsBinding {
	   usernamePassword('GOGS_USERPASS', 'gogs')
	  }
	 }
	
	 steps {
	  shell("git clone --mirror https://github.com/$gitHubUser/$gitHubRepository repo")
	  shell("cd repo && git push --mirror http://\$GOGS_USERPASS@gogs:3000/$gogsUser/$gogsRepository")
	 }
	
	 publishers {
	  mailer {
	   recipients('$ADMIN_EMAIL')
	   notifyEveryUnstableBuild(true)
	   sendToIndividuals(false)
	  }
	  wsCleanup()
	 }
	}`
)
