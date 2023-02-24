package handlers

import (
	"consumer-email-manager/services/crypto"
	"context"
	"encoding/json"

	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/streadway/amqp"
)

// Email is model for each email in app.
type Email struct {
	To      string
	Subject string
	Body    string
}

// Consume Consumes from queue then publishes messages.
func Consume(ch *amqp.Channel, client *sesv2.Client, queueName string) {
	msgs, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	for msg := range msgs {
		log.Printf("Received message %s \n", msg.Body)

		decrytypedMessage, err := crypto.Decrypt(string(msg.Body))

		if err != nil {
			log.Fatalln(err)
		}

		email := Email{}

		if err := json.Unmarshal([]byte(decrytypedMessage), &email); err != nil {
			log.Fatalln(err)
		}

		sendToSES(email, client)

		if err := msg.Ack(true); err != nil {
			log.Fatalln(err)
		}

		log.Printf("Acked message %s \n", msg.Body)
	}
}

func sendToSES(email Email, client *sesv2.Client) {
	from := os.Getenv("AWS_EMAIL_FROM")
	mailTo := email.To
	charset := aws.String("UTF-8")
	mail := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from),
		Destination: &types.Destination{
			ToAddresses: []string{mailTo},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Charset: charset,
					Data:    aws.String(email.Subject),
				},
				Body: &types.Body{
					Text: &types.Content{
						Charset: charset,
						Data:    aws.String(email.Body),
					},
				},
			},
		},
	}

	if _, err := client.SendEmail(context.Background(), mail); err != nil {
		log.Println(err)
	}

	log.Printf("Email %s sent to %s \n", email.Body, mailTo)
}
