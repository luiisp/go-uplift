package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	envs := godotenv.Load()
	if envs == nil {
		log.Fatal("Error in load envs")
	}

}
