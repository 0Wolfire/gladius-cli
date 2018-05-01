package main

import (
	"github.com/gladiusio/gladius-cli/commands"
	"github.com/gladiusio/gladius-utils/config"
)

// execute the command the user typed
func main() {
	// Setup config handling
	config.SetupConfig("test", config.CLIDefaults())
	commands.Execute()
}
