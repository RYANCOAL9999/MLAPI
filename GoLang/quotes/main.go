package main

import (
	"context"
	"flag"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ginLambda *ginadapter.GinLambda

var (
	addr_50051 = flag.String("addr", "localhost:50051", "the address to connect to")
)

func init() {
	log.Printf("Gin cold start")
	conn, err := grpc.Dial(*addr_50051, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	r := gin.Default()
	r.GET("/api/head", func(c *gin.Context) { headEvent(c, conn) })
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
