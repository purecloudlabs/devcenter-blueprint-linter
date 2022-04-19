# devcenter-blueprint-linter

A CLI tool for checking if a devcenter blueprint meets basic standards

The linter can check for remote or local repositories. The tool also supports Golang templates for customized output.

## Building the CLI

In the project root, run the command:

```bash
make build-blueprint-linter
```

This will build the binary `./bin/gc-linter` for the current OS.

## Running the Linter

To check all possible options, check the help option:

```bash
gc-linter -h
```

## Example Usage

```bash
gc_linter https://github.com/GenesysCloudBlueprints/automated-callback-blueprint -r -c ./blueprint.rule.json > result.json
```
