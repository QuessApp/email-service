package handlers

import (
	"consumer-email-manager/pkg/crypto"
	"consumer-email-manager/pkg/entities"
	"context"
	"encoding/json"

	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/streadway/amqp"
)

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
			return
		}

		email := entities.Email{}

		if err := json.Unmarshal([]byte(decrytypedMessage), &email); err != nil {
			log.Fatalln(err)
			return
		}

		if err := sendToSES(email, client); err != nil {
			log.Fatalln(err)
			return
		}

		if err := msg.Ack(true); err != nil {
			log.Fatalln(err)
			return
		}

		log.Printf("Acked message %s \n", msg.Body)
	}
}

func sendToSES(email entities.Email, client *sesv2.Client) error {
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

	_, err := client.SendEmail(context.Background(), mail)

	log.Printf("Email %s sent to %s \n", email.Body, mailTo)

	return err
}
