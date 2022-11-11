package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Create queue connection. Returns connection and channel
func Init(connUri string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(connUri)
  
  if err != nil {
    panic(err)
  }

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully connected to RabbitMQ instance at %s \n", connUri)
	}


	return conn, ch
}