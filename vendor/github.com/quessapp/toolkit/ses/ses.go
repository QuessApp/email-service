package ses

import (
	"context"

	"github.com/quessapp/toolkit/entities"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/aws-sdk-go/aws"
)

// Configure configures AWS SES client.
func Configure(accessKey, secretKey, region string) (*sesv2.Client, error) {
	amazonConfiguration, err :=
		config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion(region),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					accessKey, secretKey, "",
				),
			),
		)

	return sesv2.NewFromConfig(amazonConfiguration), err
}

// SendToSES sends an email through SES.
func SendToSES(email entities.Email, from string, client *sesv2.Client) error {
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

	return err
}
