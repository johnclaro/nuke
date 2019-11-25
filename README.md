# nuke
Removes all Docker configs, containers, images, networks, nodes, plugins, secrets, services and volumes.

## Installation

Nuke CLI is available for macOS.

### macOS

Nuke CLI is available on macOS via [Homebrew](https://brew.sh/):

TODO

## Usage

Installing the CLI provides access to the `nuke` command.

```sh-session
nuke [command]

# Run `--help` for detailed information about CLI commands
nuke [command] help
```

## Why not use go.mod?

Because, github.com/docker/docker is actually outdated.

- https://github.com/moby/moby/issues/39302
- https://www.reddit.com/r/AskProgramming/comments/atmoqa/cant_compile_sample_golang_code_from_docker/
