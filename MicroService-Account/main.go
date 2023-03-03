package main

import (
	server "account-ms/infra/api"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	time.Local = loc

	switch os.Getenv("ENVIROMENT") {
	case "local":
		log.SetOutput(os.Stdout)
	case "server":
		log.SetOutput(file)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	server.InitServer()
}
