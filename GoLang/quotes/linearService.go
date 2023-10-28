package main

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type linearGRPCClient struct {
	client pb.LinearServiceClient
}

func newLinearGRPCClient(addr string) (*linearGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &linearGRPCClient{
		client: pb.NewLinearServiceClient(conn),
	}, nil
}

func linearRegressionResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "linear regression"

		var data pb.LinearRegressionRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LinearRegressionEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}
}

func ridgeResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "ridge"

		var data pb.LinearRidgeRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LinearRidgeEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)
	}

}

func ridgeCVResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "ridge cv"

		var data pb.LinearRidgeCVRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LinearRidgeCVEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func lassoResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "lasso"

		var data pb.LassoExpressionRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LassoExpressionEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)
	}

}

func lassoLarsResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "lasso lars"

		var data pb.LassoLarsLassoExpressionRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LassoLarsLassoExpressionEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func bayesianRidgeResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "bayesian ridge"

		var data pb.BayesianRidgeRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.BayesianRidgeEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func tweedieRegressorResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "tweedie regressor"

		var data pb.TweedieRegressorRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.TweedieRegressorEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func sgdClassifierResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "sgd classifier"

		var data pb.SGDClassifierRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.SGDClassifierEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}

func elasticNetResponse(linearClient *linearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "elastic net"

		var data pb.ElasticNetRequest
		if err := readRequestData(c, &data); err != nil {
			handleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.ElasticNetEvent(ctx, &data)

		if err != nil {
			handleGRPCError(c, err, event_name)
			return
		}

		sendResponse(c, response, event_name)

	}

}
