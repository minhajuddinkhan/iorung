package main

import (
	"log"

	"github.com/minhajuddinkhan/iorung/config"
	"github.com/minhajuddinkhan/iorung/server"
)

func main() {

	conf := config.New()
	if err := server.Start(conf); err != nil {
		log.Fatal(err)
	}
}
