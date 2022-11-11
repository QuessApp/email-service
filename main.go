package main

import (
	"email-service/handlers"
	"email-service/helpers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
  
  queueUri := os.Getenv("RABBITMQ_URI")

  conn, ch := queue.Init(queueUri)

  emails.Send(ch)

  defer conn.Close()
	defer ch.Close()
}