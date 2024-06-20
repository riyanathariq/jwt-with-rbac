package main

import (
	"github.com/joho/godotenv"
	"github.com/riyanathariq/jwt-with-rbac/cmd"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided")
	}

	cmd.Start()
}
