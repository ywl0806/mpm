# MPM

## Introduction

This CLI simplifies the process of running projects from an Integrated Development Environment (IDE) or a text editor. By using this CLI, you can smoothly start a specified project by executing the command code [project name].

## Why

그냥

## Installation

```bash
go install github.com/ywl0806/mpm
```

## Usage

You can add `export PATH="$HOME/go/bin:$PATH"` to .bashrc or .zshrc.

Navigate to your project directory and execute the following command, a prompt will start.

```bash
mpm add
```

Set the project name and its corresponding directories, and input the commands and options for each directory.

```
Project Name: my-project
Choose derectories front, server
Commands for directory [front]: code
Options for directory [front]:
Commands for directory [server]: code
Options for directory [server]:
```

Use the -d or --deep option to recursively fetch directories.

```bash
# Fetch directories up to 2 levels deep from the current directory
mpm add -d 2
```

Execute the following command to run the most recently executed or added project.

```bash
mpm
```

Execute the following command to select and run a project through prompts.

```bash
mpm e
# or
mpm excure
# or
mpm run
```

```bash
# List of currently registered projects
mpm list
# Detailed information about currently registered projects
mpm list -d
```

```bash
# Select and delete a project
mpm delete
```
