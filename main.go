package main

import (
	"log"
	"os"

	"github.com/Tanay-Verma/gator/internal/command"
	"github.com/Tanay-Verma/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error: error reading config: %v", err)
	}

	state := command.NewState(&cfg)

	commands := command.NewCommands()
	commands.Register("login", command.HandlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := command.NewCommand(commandName, commandArgs)

	err = commands.Run(&state, cmd)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
