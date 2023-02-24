package main

import (
	"consumer-email-manager/handlers"
	"consumer-email-manager/pkg/queue"
	"consumer-email-manager/pkg/ses"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file")
	}

	queueURI := os.Getenv("RABBITMQ_URI")
	queueName := os.Getenv("QUEUE_NAME")

	conn, ch := queue.Connect(queueURI)
	mailClient := ses.Configure()

	handlers.Consume(ch, mailClient, queueName)

	defer conn.Close()
	defer ch.Close()
}
