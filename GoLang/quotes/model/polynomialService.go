package model

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type polynomialGRPCClient struct {
	client pb.PolynomialServiceClient
}

func newPolynomialGRPCClient(addr string) (*polynomialGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &polynomialGRPCClient{
		client: pb.NewPolynomialServiceClient(conn),
	}, nil
}

func polynomialFeaturesResponse(polynomialClient *polynomialGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "polynomial features"

		var data pb.PolynomialFeaturesRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := polynomialClient.client.PolynomialFeaturesEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func polynomialFeaturesFitTransformResponse(polynomialClient *polynomialGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "polynomial features fit transform"

		var data pb.PolynomialFeaturesFitTransformRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := polynomialClient.client.PolynomialFeaturesFitTransformEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)
	}

}
