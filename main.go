package main

import (
	"fmt"
	"log"

	"github.com/Tanay-Verma/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("tanay")
	if err != nil {
		log.Fatalf("error setting the user in config: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println(cfg)
}
