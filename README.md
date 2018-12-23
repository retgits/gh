# gh - a collection of git helper commands

[![Travis](https://img.shields.io/travis/retgits/gh.svg?style=flat-square)](https://travis-ci.org/retgits/gh)
[![license](https://img.shields.io/github/license/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/releases)

![gh](./gh.png)

A collection of git helper commands to make my life a little easier

## install

```bash
go get -u github.com/retgits/gh
```

## Usage

```bash
Usage:
  gh [command]

Available Commands:
  all         Stage all unstaged files
  amend       Use the last commit message and amend your stuffs
  clone       Clone a repository to a specified directory
  commit      A simpler alias for "git commit -a -S -m"
  credit      A very slightly quicker way to credit an author on the latest commit
  dash        Update the snippets in Dash with GitHub gists
  git         Create GitHub and/or Gogs repositories and optionally a Jenkins job as well.
  help        Help about any command
  lambda      Create a new AWS Lambda function based on my personal templates in the current folder.
  nuke        Removes a branch locally and on the remote origin
  travis      Update the AWS credentials on Travis-CI jobs.
  undo        Undo the last commit, but don't throw away any changes

Flags:
  -h, --help      help for gh
      --version   version for gh

Use "gh [command] --help" for more information about a command.
```

### All

```bash
Stage all unstaged files

Usage:
  gh all [flags]

Flags:
  -h, --help   help for all
```

### Amend

```bash
Use the last commit message and amend your stuffs

Usage:
  gh amend [flags]

Flags:
  -h, --help   help for amend
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

### Commit

```bash
A simpler alias for "git commit -a -S -m"

Usage:
  gh commit [flags]

Flags:
  -h, --help             help for commit
      --message string   The commit message (required)
```

### Credit

```bash

```

### Dash

```bash
a command to update the snippets in Dash with GitHub gists.

Usage:
  gh dash [flags]

Flags:
      --github-token string   The Personal Access Token for GitHub (optional)
  -h, --help                  help for dash
      --lib string            The full path to the library.dash file (like /Users/username/Library/Application Support/Dash/library.dash)
      --owner string          The GitHub username to get gists for (required)
```

For the GitHub token. The precedence is as follows:

* Flag   : github-token
* Env var: GITHUBTOKEN

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

### Nuke

```bash
Removes a branch locally and on the remote origin

Usage:
  gh nuke [flags]

Flags:
      --branch string   The Nuke message (required)
  -h, --help            help for nuke
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

## License

See the [LICENSE](./LICENSE) file in the repository