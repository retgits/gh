# all

```bash
Stage all unstaged files

Usage:
  gh all [flags]

Flags:
  -h, --help   help for all

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# amend

```bash
Use the last commit message and amend your stuffs

Usage:
  gh amend [flags]

Flags:
  -h, --help   help for amend

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# clone

```bash
clone makes sure repositories are cloned to a specified base directory and a predefined structure:
  <basefolder>/<git site>/<user>/<repo> (like /home/user/github.com/retgits/gh). The basefolder is
  set by the git.basefolder in .ghconfig.yml or the --basefolder flag

sample usage: gh clone https://github.com/retgits/gh

Usage:
  gh clone [flags]

Flags:
      --basefolder string   The root folder to clone to (this flag overrides git.basefolder from the configuration file)
  -h, --help                help for clone

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# commit

```bash
A simpler alias for "git commit -a -S -m"

Usage:
  gh commit [flags]

Flags:
  -h, --help             help for commit
      --message string   The commit message (required)

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# credit

```bash
A very slightly quicker way to credit an author on the latest commit

Usage:
  gh credit [flags]

Flags:
      --email string   The email address of the author to credit (required)
  -h, --help           help for credit
      --name string    The name of the author to credit (required)

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# github

```bash
Create a GitHub repository

Usage:
  gh github [flags]

Flags:
      --ghrepo string    The repository name to create (will default to the name of the directory if not set)
      --ghtoken string   The Personal Access Token for GitHub (this flag overrides git.ghtoken from the configuration file)
  -h, --help             help for github

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# gogs

```bash
Create a Gogs repository

Usage:
  gh gogs [flags]

Flags:
      --gogsrepo string    The repository name to create (will default to the name of the directory if not set)
      --gogstoken string   The Personal Access Token for gogs (this flag overrides git.gogstoken from the configuration file)
      --gogsurl string     The URL of the gogs server (this flag overrides git.gogsurl from the configuration file)
  -h, --help               help for gogs

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# jenkins

```bash
Create a Jenkins Job

Usage:
  gh jenkins [flags]

Flags:
      --commit               Commit and push the updates to the Jenkins DSL project
  -h, --help                 help for jenkins
      --jenkinsrepo string   The location of the Jenkins Job DSL project (this flag overrides git.jenkinsrepo from the configuration file)
      --projectname string   The name of the project to create a new job for (will default to the name of the directory if not set)

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# nuke

```bash
Removes a branch locally and on the remote origin

Usage:
  gh nuke [flags]

Flags:
      --branch string   The Nuke message (required)
  -h, --help            help for nuke

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

# undo

```bash
Undo the last commit, but don't throw away any changes

Usage:
  gh undo [flags]

Flags:
  -h, --help   help for undo

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```