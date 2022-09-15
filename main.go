package main

import (
	server "go-api/infra/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	switch os.Getenv("ENVIROMENT") {
	case "local":
		err = godotenv.Load(".env.local")
		log.SetOutput(os.Stdout)
	case "production":
		err = godotenv.Load(".env.production")
		log.SetOutput(file)
	case "development":
		err = godotenv.Load(".env.development")
		log.SetOutput(file)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	server.InitServer()
}
