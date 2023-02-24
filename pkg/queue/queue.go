package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// Connect connects to RabbitMQ.
func Connect(connUri string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(connUri)

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	log.Printf("Successfully connected to RabbitMQ instance at %s \n", connUri)

	return conn, ch
}
