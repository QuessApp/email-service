package emails

import (
	"context"
	"email-service/services/crypto"
	"encoding/json"
	"fmt"
	"os"

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
    fmt.Printf("Received message %s \n", msg.Body)
    
    decrytypedMessage, decryptError := crypto.Decrypt(string(msg.Body))

    if decryptError != nil {
      panic(decryptError)
    }
    
    var emailPayload Email 

    if unmarshEmailPayloadError := json.Unmarshal([]byte(decrytypedMessage), &emailPayload); unmarshEmailPayloadError != nil {
      fmt.Print(unmarshEmailPayloadError)
    }

    sendToAwsSES(emailPayload, mailClient)

    ackMessageError := msg.Ack(true)

    if ackMessageError != nil {
      fmt.Print(ackMessageError)
    }

    fmt.Printf("Acked message %s \n", msg.Body)
  }
}

func sendToAwsSES(emailPayload Email, mailClient *sesv2.Client) {
  from := os.Getenv("AWS_EMAIL_FROM")
  mailTo := emailPayload.To
  charset := aws.String("UTF-8")
  mail := &sesv2.SendEmailInput{
    FromEmailAddress: aws.String(from),
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
    fmt.Print(sendEmailError)
  }

  fmt.Printf("Email %s sent to %s \n", emailPayload.Body, mailTo)
}