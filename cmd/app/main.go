package main

import (
	"email-service/internal/consumer"
	"log"
	"os"

	"github.com/quessapp/toolkit/queue"
	"github.com/quessapp/toolkit/ses"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file")
	}

	queueURI := os.Getenv("RABBITMQ_URI")
	queueName := os.Getenv("QUEUE_NAME")
	cipherKey := os.Getenv("CIPHER_KEY")

	mailFrom := os.Getenv("AWS_EMAIL_FROM")
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	region := os.Getenv("AWS_REGION")

	conn, ch := queue.Connect(queueURI)
	mailClient, err := ses.Configure(accessKey, secretKey, region)

	if err != nil {
		panic(err)
	}

	consumer.Consume(ch, mailClient, queueName, cipherKey, mailFrom)

	defer conn.Close()
	defer ch.Close()
}
