package consumer

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/quessapp/toolkit/crypto"
	"github.com/quessapp/toolkit/entities"
	"github.com/quessapp/toolkit/queue"
	"github.com/quessapp/toolkit/ses"
	"github.com/streadway/amqp"
)

// Consume Consumes from queue then publishes messages.
func Consume(ch *amqp.Channel, client *sesv2.Client, queueName, cipherKey, mailFrom string) {
	msgs, err := queue.Consume(ch, queueName)

	if err != nil {
		log.Fatalln(err)
	}

	for msg := range msgs {
		log.Printf("Received message %s \n", msg.Body)

		decrytypedMessage, err := crypto.Decrypt(string(msg.Body), cipherKey)

		if err != nil {
			log.Fatalln(err)
			return
		}

		email := entities.Email{}

		if err := json.Unmarshal([]byte(decrytypedMessage), &email); err != nil {
			log.Fatalf("Error unmarshalling message: %s \n", err)
			return
		}

		if err := ses.SendToSES(email, mailFrom, client); err != nil {
			log.Println(err)
			return
		}

		if err := msg.Ack(true); err != nil {
			log.Fatalln(err)
			return
		}

		log.Printf("Email from %s was sent to %s! \n", mailFrom, email.To)
		log.Printf("Acked message %s \n", msg.Body)
	}
}
