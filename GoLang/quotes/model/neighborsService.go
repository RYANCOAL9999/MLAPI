package model

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NeighborsGPRCClient struct {
	client pb.NeighborsServiceClient
}

func NewNeighborsGPRCClient(addr string) (*NeighborsGPRCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &NeighborsGPRCClient{
		client: pb.NewNeighborsServiceClient(conn),
	}, nil
}

func NearestNeighborsResponse(neighborsClient *NeighborsGPRCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "nearest neighbors"

		var data pb.NearestNeighborsRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := neighborsClient.client.NearestNeighborsEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func KDTreeResponse(neighborsClient *NeighborsGPRCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "kd tree"

		var data pb.KDTreeRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := neighborsClient.client.KDTreeEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func NearestCentroidResponse(neighborsClient *NeighborsGPRCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "nearest centroid"

		var data pb.NearestCentroidRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := neighborsClient.client.NearestCentroidEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}
