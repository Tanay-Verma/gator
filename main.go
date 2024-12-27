package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Tanay-Verma/gator/internal/command"
	"github.com/Tanay-Verma/gator/internal/config"
	"github.com/Tanay-Verma/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error: error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)

	state := command.NewState(&cfg, dbQueries)

	commands := command.NewCommands()
	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerUsers)
	commands.Register("agg", command.HandlerAgg)
	commands.Register("addfeed", command.HandlerAddFeed)
	commands.Register("feeds", command.HandlerFeeds)
	commands.Register("follow", command.HandlerFollow)
	commands.Register("following", command.HandlerFollowing)

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
