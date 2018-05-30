# ghhelper - git helper
a githelper command to create a GitHub and/or Gogs repository and optionally a Jenkins job as well. Available flags:

* github: A boolean flag to create a GitHub repository for this project (defaults to false)
* gogs: A boolean flag to create a Gogs repository for this project (defaults to false)
* jenkins: A boolean flag to create a Jenkins DSL for this project (defaults to false)
* commit: A boolean flag to commit and push the updates to the Jenkins DSL project (defaults to false)
* github-token: The token to use to connect to GitHub (optional)
* gogs-token: The token to use to connect to Gogs (optional)
* jenkins-base: The base directory of the Jenkins DSL project (optional)

## usage
```
$ ghhelper [--github] [--gogs] [--jenkins] [--commit] [--github-token token] [--gogs-token token] [--jenkins-base dir]
```

For the GitHub token. The precedence is as follows:

* Flag   : github-token
* Env var: GITHUBTOKEN

For the Gogs token. The precedence is as follows:

* Flag   : gogs-token
* Env var: GOGSTOKEN

For the Jenkins directory.The precedence is as follows:

* Flag   : jenkins-base
* Env var: JENKINSBASEDIR

## install
```
$ go get -u github.com/retgits/ghhelper
```