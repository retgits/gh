# gh - a collection of git helper commands

[![Travis](https://img.shields.io/travis/retgits/gh.svg?style=flat-square)](https://travis-ci.org/retgits/gh)
[![License](https://img.shields.io/github/license/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/retgits/gh)](https://goreportcard.com/report/github.com/retgits/gh)
[![Stability](https://img.shields.io/badge/stability-stable-green.svg?style=flat-square)](https://img.shields.io/badge/stability-stable-green.svg?style=flat-square)

![gh](./gh.png)

`gh` is a collection of git helper commands to make my life a little easier. The command-line tool wraps a number of git commands that I frequently use.

## Install

To install `gh` from source, run

```bash
go get -u github.com/retgits/gh
```

Or get a release version from the [releases](./releases) tab.

## Configuration

Configuration is done using flags for the commands, or using the `.ghconfig.yml` file in your `HOME` directory. The config file has the following parameters:

```yml
git:
  basefolder:  ## The base folder to clone repositories to
github:
  accesstoken: ## The personal access token to connect to GitHub
gogs:
  accesstoken: ## The personal access token to connect to Gogs
  apiendpoint: ## The API endpoint  of the Gogs server (like http://localhost/api/v1)
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
  create-repo Create a repository
  credit      A very slightly quicker way to credit an author on the latest commit
  help        Help about any command
  nuke-branch Removes a branch locally and on the remote origin
  undo        Undo the last commit, but don't throw away any changes

Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
  -h, --help            help for gh
      --version         version for gh

Use "gh [command] --help" for more information about a command.
```

For more detailed information on the commands, check the [docs](./docs/commands.md)

## License

See the [LICENSE](./LICENSE) file in the repository