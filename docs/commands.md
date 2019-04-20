# Commands

## all

```bash
Stage all unstaged files

Usage:
  gh all [flags]

Flags:
  -h, --help   help for all

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

## amend

```bash
Use the last commit message and amend your stuffs

Usage:
  gh amend [flags]

Flags:
  -h, --help   help for amend

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

## clone

```bash
clone makes sure repositories are cloned to a specified base directory and a predefined structure:
  <basefolder>/<git site>/<user>/<repo> (like /home/user/github.com/retgits/gh). The basefolder is
  set by the git.basefolder in .ghconfig.yml or the --basefolder flag

sample usage: gh clone https://github.com/retgits/gh

Usage:
  gh clone [flags]

Flags:
      --basefolder string   The root folder to clone to
  -h, --help                help for clone

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

## commit

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

## create-repo

```bash
Create a repository

Usage:
  gh create-repo [flags]

Flags:
  -h, --help           help for create-repo
      --org string     The organization to create the repo under
      --private        Set to true to create a private repository
      --repo string    The repository name to create
      --token string   The Personal Access Token for the version control system
      --type string    The version control system to use
      --url string     The API endpoint to call

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

## credit

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

## nuke-branch

```bash
Removes a branch locally and on the remote origin

Usage:
  gh nuke-branch [flags]

Flags:
      --branch string   The branch to remove (required)
  -h, --help            help for nuke-branch

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```

## undo

```bash
Undo the last commit, but don't throw away any changes

Usage:
  gh undo [flags]

Flags:
  -h, --help   help for undo

Global Flags:
      --config string   config file (default is $HOME/.ghconfig.yml)
```