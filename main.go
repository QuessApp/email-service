package main

import (
	"email-service/handlers"
	"email-service/services/queue"
	"email-service/services/ses"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  loadEnvFileError := godotenv.Load(".env")

  if loadEnvFileError != nil {
    log.Fatalf("Error loading .env file")
  }
  
  queueUri := os.Getenv("RABBITMQ_URI")

  conn, ch := queue.Init(queueUri)
  mailClient := ses.Init()
  
  emails.Send(ch, mailClient)

  defer conn.Close()
	defer ch.Close()
}