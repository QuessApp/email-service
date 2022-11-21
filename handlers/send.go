package emails

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/streadway/amqp"
)

type Email struct {
  To string 
  Subject string
  Body string 
}

// Consumes from queue then publishes messages
func Send(ch *amqp.Channel, mailClient *sesv2.Client) {
	msgs, consumeError := ch.Consume(
    "SendEmail", 
    "",
    false,
    false,
    false,
    false,
    nil,
  )

  if consumeError != nil {
    panic(consumeError)
  }
  
  for msg := range(msgs) {
    ackMessageError := msg.Ack(true)

    if ackMessageError != nil {
      log.Fatal(ackMessageError)
    }

    fmt.Printf("Received and acked message %s \n", msg.Body)

    var emailPayload Email 

    if unmarshEmailPayloadError := json.Unmarshal(msg.Body, &emailPayload); unmarshEmailPayloadError != nil {
      log.Fatal(unmarshEmailPayloadError)
    }

    sendToAwsSES(emailPayload, mailClient)
  }
}

func sendToAwsSES(emailPayload Email, mailClient *sesv2.Client) {
  mailTo := emailPayload.To
  charset := aws.String("UTF-8")
  mail := &sesv2.SendEmailInput{
    FromEmailAddress: aws.String(mailTo),
    Destination: &types.Destination{
      ToAddresses: []string{ mailTo },
    },
    Content: &types.EmailContent{
      Simple: &types.Message{
        Subject: &types.Content{
          Charset: charset,
          Data: aws.String(emailPayload.Subject),
        },
        Body: &types.Body{
          Text: &types.Content{
            Charset: charset,
            Data: aws.String(emailPayload.Body),
          },
        },
      },
    },
  }

  // TODO: RECOVER APP WHEN EMAIL DOES NOT EXISTS
  _, sendEmailError := mailClient.SendEmail(context.Background(), mail)

  if sendEmailError != nil {
    log.Fatal(sendEmailError)
  }

  fmt.Printf("Email %s sent to %s \n", emailPayload.Body, mailTo)
}