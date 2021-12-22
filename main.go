package main

import (
	"log"

	"hc21f/pkg/database"
	"hc21f/pkg/twitter"
	"hc21f/runner"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := twitter.Init(); err != nil {
		log.Fatal(err)
	}

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}

	t := twitter.Get()
	db := database.Get()

	runner.Run(t, db)
}
