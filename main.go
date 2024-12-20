package main

import (
	"fmt"

	"github.com/Tanay-Verma/gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("tanay")
	cfg = config.Read()
	fmt.Println(cfg)
}
