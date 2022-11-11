package emails

import (
	"log"
	"github.com/streadway/amqp"
)

// Consumes from queue then publishes messages
func Send(ch *amqp.Channel) {
	msgs, err := ch.Consume(
    "SendEmailReceivedNewQuestion", 
    "",
    false,
    false,
    false,
    false,
    nil,
  )

  if err != nil {
    panic(err)
  }
  
  for msg := range(msgs) {
    err := msg.Ack(true)

    if err != nil {
      log.Fatal(err)
    }

    log.Printf("Received and acked a message: %s", msg.Body)
  }
}