package main

import (
	"context"
	"net/http"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type svmGRPCClient struct {
	client pb.SVMServiceClient
}

func newSVMGRPCClient(addr string) (*svmGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &svmGRPCClient{
		client: pb.NewSVMServiceClient(conn),
	}, nil
}

func linearSVCResponse(svmClient *svmGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "linear svc"

		var data pb.LinearRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, http.StatusBadRequest, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := svmClient.client.LinearSVCEvent(ctx, &data)
		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)
	}
}

func linearSVRResponse(svmClient *svmGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "linear svr"

		var data pb.LinearRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, http.StatusBadRequest, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := svmClient.client.LinearSVREvent(ctx, &data)
		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)
	}
}

func svcResponse(svmClient *svmGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "svc"

		var data pb.SVCRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, http.StatusBadRequest, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := svmClient.client.SVCEvent(ctx, &data)
		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)
	}
}
