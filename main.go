/*
tag generates artwork from your favourite terminal themes.

Usage:

	tag [command]

Available Commands:

	completion  Generate the autocompletion script for the specified shell
	generate    Create a new generative artwork
	help        Help about any command

Flags:

	-h, --help      help for tag
	-v, --version   version for tag

Use "tag [command] --help" for more information about a command.
*/
package main

import "github.com/rsHalford/tag/cmd"

func main() {
	cmd.Execute()
}
