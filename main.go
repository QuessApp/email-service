package main

import (
	emails "consumer-email-manager/handlers"
	"consumer-email-manager/services/queue"
	"consumer-email-manager/services/ses"
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