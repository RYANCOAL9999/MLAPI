package main

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SVMGRPCClient struct {
	client pb.SVMServiceClient
}

func NewSVMGRPCClient(addr string) (*SVMGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &SVMGRPCClient{
		client: pb.NewSVMServiceClient(conn),
	}, nil
}

func LinearSVCResponse(svmClient *SVMGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "linear svc"

		var data pb.LinearRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := svmClient.client.LinearSVCEvent(ctx, &data)
		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)
	}
}

func LinearSVRResponse(svmClient *SVMGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "linear svr"

		var data pb.LinearRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := svmClient.client.LinearSVREvent(ctx, &data)
		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)
	}
}

func SVCResponse(svmClient *SVMGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "svc"

		var data pb.SVCRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := svmClient.client.SVCEvent(ctx, &data)
		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)
	}
}
