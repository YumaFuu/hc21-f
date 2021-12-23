package main

import (
	"log"

	"hc21f/pkg/runner"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	runner.Run()
}
