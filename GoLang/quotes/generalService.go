package main

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type generalGRPCClient struct {
	client pb.GeneralServiceClient
}

func newGeneralGRPCClient(addr string) (*generalGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &generalGRPCClient{
		client: pb.NewGeneralServiceClient(conn),
	}, nil
}

func infoResponse(generalClient *generalGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "info"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := generalClient.client.InfoEvent(ctx, &pb.InfoRequest{})

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func describleResponse(generalGRPCClient *generalGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "describle"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := generalGRPCClient.client.DescriblerEvent(ctx, &pb.DescribeRequest{})

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func headerResponse(generalGRPCClient *generalGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "header"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := generalGRPCClient.client.HeaderEvent(ctx, &pb.HeaderRequest{})

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}
}
