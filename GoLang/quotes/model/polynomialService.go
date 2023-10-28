package main

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PolynomialGRPCClient struct {
	client pb.PolynomialServiceClient
}

func NewPolynomialGRPCClient(addr string) (*PolynomialGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &PolynomialGRPCClient{
		client: pb.NewPolynomialServiceClient(conn),
	}, nil
}

func PolynomialFeaturesResponse(polynomialClient *PolynomialGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "polynomial features"

		var data pb.PolynomialFeaturesRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := polynomialClient.client.PolynomialFeaturesEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func PolynomialFeaturesFitTransformResponse(polynomialClient *PolynomialGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "polynomial features fit transform"

		var data pb.PolynomialFeaturesFitTransformRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := polynomialClient.client.PolynomialFeaturesFitTransformEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)
	}

}
