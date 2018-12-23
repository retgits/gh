# gh - a collection of git helper commands

[![Travis](https://img.shields.io/travis/retgits/gh.svg?style=flat-square)](https://travis-ci.org/retgits/gh)
[![license](https://img.shields.io/github/license/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/retgits/gh.svg?style=flat-square)](https://github.com/retgits/gh/releases)

![gh](./gh.png)

A collection of git helper commands to make my life a little easier

## Install

```bash
go get -u github.com/retgits/gh
```

## Configuration

Configuration is done using flags for the commands, or using the `.ghconfig.yml` file in your `HOME` directory. The config file has the following parameters:

```yml
git:
  basefolder:  ## The base folder to clone repositories to
  ghtoken:     ## The personal access token to connect to GitHub
  gogstoken:   ## The personal access token to connect to Gogs
  gogsurl:     ## The URL of the Gogs server
  jenkinsrepo: ## The location of the Jenkins Job DSL repo on disk
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
  github      Create a GitHub repository
  gogs        Create a Gogs repository
  help        Help about any command
  jenkins     Create a Jenkins Job
  nuke        Removes a branch locally and on the remote origin
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