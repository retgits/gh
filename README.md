# gh - a collection of git helper commands

[![Travis](https://img.shields.io/travis/retgits/gh.svg?style=flat-square)](https://travis-ci.org/retgits/gh)
[![license](https://img.shields.io/github/license/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/releases)

A collection of git helper commands to make my life a little easier

## install

```bash
$ go get -u github.com/retgits/gh
```

## Usage

```bash
Usage:
  gh [command]

Available Commands:
  clone       a simple git clone command to make sure that all git clones end up in a specified directory.
  git         a git helper command to create a GitHub and/or Gogs repository and optionally a Jenkins job as well.
  help        Help about any command
  lambda      a command to create a new AWS Lambda function based on my personal templates in the current folder.
  travis      a command to update the AWS credentials on Travis-CI jobs.

Flags:
  -h, --help      help for gh
      --version   version for gh

Use "gh [command] --help" for more information about a command.
```

### Clone

```bash
gh clone is a simple git clone command to make sure that all git clones end up in a specified
directory. The directory is specified by
1) setting a flag `base` (gh clone --base . https://github.com/retgits/gh)
2) setting an environment variable `GITBASEFOLDER`
3) the current directory

Sample usage: gh clone https://github.com/retgits/gh

Usage:
  gh clone [flags]

Flags:
      --base string   The root folder to clone this repo in (optional, unless $GITBASEFOLDER is set)
  -h, --help          help for clone
```

### Git

```bash
a git helper command to create a GitHub and/or Gogs repository and optionally a Jenkins job as well.

Usage:
  gh git [flags]

Flags:
      --commit                Commit and push the updates to the Jenkins DSL project
      --github                Create a GitHub repository for this project
      --github-token string   The Personal Access Token for GitHub (optional)
      --gogs                  Create a Gogs repository for this project
      --gogs-token string     The Personal Access Token for Gogs (optional)
  -h, --help                  help for git
      --jenkins               Create a Jenkins DSL for this project
      --jenkins-base string   The base directory of the Jenkins DSL project (optional)
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

### Lambda

```bash
gh lambda is a command to create a new AWS Lambda function based on my personal templates in the current folder

Sample usage: gh lambda my-lambda
This will create a new AWS Lambda function in the my-lambda folder of this directory

Usage:
  gh lambda [flags]

Flags:
      --base string   The root folder to create this lambda function in (optional, will default to current folder)
  -h, --help          help for lambda
      --name string   The name of the lambda function you want to create (required)
```

### Travis

```bash
a command to update the AWS credentials on Travis-CI jobs.

Usage:
  gh travis [flags]

Flags:
  -h, --help                  help for travis
      --owner string          The owner of the Travis-CI repos (required)
      --repos string          The list of Travis-CI repos to update (optional, must be a comma separated list)
      --travis-token string   The Authentication Token for Travis-CI (optional)
```

For the Travis-CI token. The precedence is as follows:

* Flag   : travis-token
* Env var: TRAVISTOKEN

For the repo list. The precedence is as follows:

* Flag   : repos
* Env var: TRAVISREPOS

The new values for **AWS_ACCESS_KEY_ID** and **AWS_SECRET_ACCESS_KEY** are retrieved using the `aws configure get` command.