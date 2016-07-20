# git-user

Git plugin that allows you to save multiple user profiles and set them as project defaults

### Why?
I created this because I have multiple emails that I use for work, personal, and open source projects and I would find myself checking into a git project with the wrong profile. Instead of manually changing the git config for a project every time, this was much easier.

## Installing

### Compile
This project requires go1.6 to compile. Just run `go get -u github.com/gesquive/git-user` and the executable should be built for you automatically in your `$GOPATH`.

Optionally you can run `make install` to build and copy the executable to `/usr/local/bin/` with correct permissions.

### Download
Alternately, you can download the latest release for your platform from [github](https://github.com/gesquive/git-user/releases).

Once you have an executable, make sure to copy it somewhere on your path like `/usr/local/bin` or `C:/Program Files/`.
If on a \*nix/mac system, make sure to run `chmod +x /path/to/git-user`.

## Usage

If the `git-user` executable is placed on the path, it can be used as a git command. For example, you would be able to run the command `git user list` to list all of the configured users.


```console
git-user lets you quickly switch between multiple git user profiles

Usage:
  git-user [flags] [command]

Available Commands:
  add         Add a new profile
  del         Delete a profile
  edit        Edit a profile
  list        List all saved profiles
  rm          Remove a profile from the current project
  set         Set the profile for the current project

Flags:
  -c, --config string     config file (default "~/.git-profiles")
  -g, --git-path string   The git executable to use (default "git")
  -p, --path string       The project to get/set the user (default "$CWD")
  -V, --version           Show the version and exit
```

Optionally, a hidden debug flag is available in case you need additional output.
```console
Hidden Flags:
  -D, --debug                  Include debug statements in log output
```

## QuickStart

```console
$ cd /path/to/git/project

# add a home profile for Tobe
$ git user add home "Tobe Developer" tobe@home.com
Added profile 'test'

# add a work profile for Tobe
$ git user add work "Tobe A Developer" tobe@work.com

# list out our saved profiles
$ git user list
Global Profile:
  User: Tobe <tobe@default.com>

Saved Profiles:
  home: Tobe Developer <tobe@home.com>
  work: Tobe A Developer <tobe@work.com>

# set the current git repository user to the home profile
$ git user set home
The user for the 'project' repository has been set too 'Tobe Developer <tobe@home.com>'

# list profiles again, notice how the current repository profile is now set
$ git user
Project Profile:
  Path: /path/to/git/project
  User: Tobe Developer <tobe@home.com>

Saved Profiles:
  home: Tobe Developer <tobe@home.com>
  work: Tobe A Developer <tobe@work.com>
```


## Documentation

This documentation can be found at github.com/gesquive/git-user

## License

This package is made available under an MIT-style license. See LICENSE.

## Contributing

PRs are always welcome!


<!-- TODO: Generate a manpage -->
<!-- TODO: Include a detailed install script -->
