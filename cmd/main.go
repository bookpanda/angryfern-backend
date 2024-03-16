package main

import (
	"fmt"

	"github.com/bookpanda/angryfern-backend/cfgldr"
	"github.com/bookpanda/angryfern-backend/database"
)

func main() {
	conf, err := cfgldr.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	_, err = database.New(conf)
	if err != nil {
		panic(fmt.Sprintf("Failed to load database: %v", err))
	}
}
