package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Create queue connection. Returns connection and channel
func Init(connUri string) (*amqp.Connection, *amqp.Channel) {
	conn, connectToQueueError := amqp.Dial(connUri)
  
  if connectToQueueError != nil {
    panic(connectToQueueError)
  }

	ch, connectToChannelError := conn.Channel()

	if connectToChannelError != nil {
		panic(connectToChannelError)
	} else {
		fmt.Printf("Successfully connected to RabbitMQ instance at %s \n", connUri)
	}


	return conn, ch
}