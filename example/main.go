package main

import (
	"log"
	"os"

	"genv"
)

func init() {
	if err := genv.New(".env", ".env.example"); err != nil {
		log.Printf("error -> %v", err)
	}
}

func main() {
	log.Printf("defined env: %s\ndefined example from other .env file: %s", os.Getenv("TEST_KEY"), os.Getenv("TEST_EXAMPLE_KEY"))
}
