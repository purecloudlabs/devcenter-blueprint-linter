# devcenter-blueprint-linter

A CLI tool for checking if a devcenter blueprint meets basic standards

The linter can check for remote or local repositories. The tool also supports Golang templates for customized output.

## Building the CLI

In the project root, run the command:

```bash
make build
```

This will build the binary `./bin/gc_linter` for the current OS.

## Running the Linter

To check all possible options, check the help option:

```bash
gc_linter -h
```

## Example Usage

```bash
gc_linter https://github.com/GenesysCloudBlueprints/automated-callback-blueprint -r -c ./blueprint.rule.json > result.json
```

## Logging

To enable logging, add the flag `-l` or `--enable-logging` to the command.

The log file should appear in the following directories depending on your OS:

Windows: `%TEMP%\GenesysCloud`
Unix: `/tmp/GenesysCloud`
OSX: `<home directory>/Library/Logs/GenesysCloud`
