package main

import (
	"context"
	"log"
	"os"
	"rest-api/hello"
	"rest-api/routes"
	"rest-api/user"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

var ddb *dynamodb.Client

func main() {

	initDynamoDB()

	r := gin.Default()

	// Register routes from other packages
	routes.RegisterHelloRoute(r, InitializeHelloHandler())

	routes.RegisterUserRoute(r, InitializeUserHandler(ddb))

	if err := r.Run(":3600"); err != nil {
		log.Fatal(err)
	}
}

func InitializeHelloHandler() *hello.HelloHandler {
	return hello.NewHelloHandler()
}

func InitializeUserHandler(ddb *dynamodb.Client) *user.UserHandler {
	userService := user.NewUserService(ddb)
	return user.NewUserHandler(userService)
}

func initDynamoDB() {
	// Load base AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("test", "test", ""),
		),
	)

	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	var endpoint *string

	env := os.Getenv("APP_ENV")
	url := "http://localhost:4566"
	if env == "dev" {
		url = "http://storage:4566"
	}

	endpoint = &url

	// Construct the DynamoDB client using BaseEndpoint (this is not deprecated)
	ddb = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		if endpoint != nil {
			o.BaseEndpoint = endpoint
		}
	})
}
