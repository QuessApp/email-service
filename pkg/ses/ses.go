package ses

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

// Configure configures AWS SES client.
func Configure() *sesv2.Client {
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	region := os.Getenv("AWS_REGION")

	amazonConfiguration, createAmazonConfigurationError :=
		config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion(region),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					accessKey, secretKey, "",
				),
			),
		)

	if createAmazonConfigurationError != nil {
		log.Fatal(createAmazonConfigurationError)
	}

	return sesv2.NewFromConfig(amazonConfiguration)
}
