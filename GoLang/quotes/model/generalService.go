package model

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GeneralGRPCClient struct {
	client pb.GeneralServiceClient
}

func NewGeneralGRPCClient(addr string) (*GeneralGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GeneralGRPCClient{
		client: pb.NewGeneralServiceClient(conn),
	}, nil
}

func InfoResponse(generalClient *GeneralGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "info"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := generalClient.client.InfoEvent(ctx, &pb.InfoRequest{})

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func DescriblerResponse(generalGRPCClient *GeneralGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "describle"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := generalGRPCClient.client.DescriblerEvent(ctx, &pb.DescribeRequest{})

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func HeaderResponse(generalGRPCClient *GeneralGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "head"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := generalGRPCClient.client.HeaderEvent(ctx, &pb.HeaderRequest{})

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}
}
