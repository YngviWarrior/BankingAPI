package main

import (
	server "api-go/infra/api"
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

	switch os.Getenv("ENVIROMENT") {
	case "local":
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		time.Local = loc

		log.SetOutput(file)
		log.SetOutput(os.Stdout)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	server.InitServer()
}
