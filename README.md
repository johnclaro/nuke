# nuke
A command-line tool for removing Docker configs, containers, images, networks and volumes all at once.

## Installation

```
go install
```

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
