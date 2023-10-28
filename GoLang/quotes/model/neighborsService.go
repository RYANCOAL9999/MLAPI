package model

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type neighborsGPRCClient struct {
	client pb.NeighborsServiceClient
}

func newNeighborsGPRCClient(addr string) (*neighborsGPRCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &neighborsGPRCClient{
		client: pb.NewNeighborsServiceClient(conn),
	}, nil
}

func nearestNeighborsResponse(neighborsClient *neighborsGPRCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "nearest neighbors"

		var data pb.NearestNeighborsRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := neighborsClient.client.NearestNeighborsEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func kdTreeResponse(neighborsClient *neighborsGPRCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "kd tree"

		var data pb.KDTreeRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := neighborsClient.client.KDTreeEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func nearestCentroidResponse(neighborsClient *neighborsGPRCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "nearest centroid"

		var data pb.NearestCentroidRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := neighborsClient.client.NearestCentroidEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}
