# scratchpad

Create a scratchpad to play/prototype with your favorite programming language.
(Similar to IntelliJ's scratch files, but from the terminal.)

**Supported Languages**: Bash, Go, Java, Javascript, Lua, Python, Rust

This CLI simply executes the initialization commands for these languages (check `mappings` in `lang.go`).

##### (**Please Note**: This assumes that the necessary tooling is already installed for the programming language.)

## Motivation

-   Wanted to create a quick scratchpad file to explore some language features/prototype something from the terminal, but had to repeatedly do these tasks first -
    -   optionally create a directory
    -   type the initialization command (especially those long `java mvn` commands)
-   I wanted to write something with the [Cobra][cobra] library :smile:

## Install

```sh
go build -o ./bin/scratchpad .

# Copy the binary somewhere in the $PATH
cp ./bin/scratchpad ~/.local/bin/scratchpad
```

## Uninstall

```sh
rm -i $(which scratchpad)
```

## Usage

```sh
scratchpad create --lang java --name "Scratch_Java"
# (or)
scratchpad create -l go -n "Go_Playground"
```

[cobra]: https://github.com/spf13/cobra
