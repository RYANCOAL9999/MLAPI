package main

import (
	"context"
	"time"

	pb "github.com/RYANCOAL9999/AI_Go_Proto_File"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LinearGRPCClient struct {
	client pb.LinearServiceClient
}

func NewLinearGRPCClient(addr string) (*LinearGRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &LinearGRPCClient{
		client: pb.NewLinearServiceClient(conn),
	}, nil
}

func LinearRegressionResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "linear regression"

		var data pb.LinearRegressionRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LinearRegressionEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}
}

func RidgeResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "ridge"

		var data pb.LinearRidgeRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LinearRidgeEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)
	}

}

func RidgeCVResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "ridge cv"

		var data pb.LinearRidgeCVRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LinearRidgeCVEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func LassoResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "lasso"

		var data pb.LassoExpressionRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LassoExpressionEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)
	}

}

func LassoLarsResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "lasso lars"

		var data pb.LassoLarsLassoExpressionRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.LassoLarsLassoExpressionEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func BayesianRidgeResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "bayesian ridge"

		var data pb.BayesianRidgeRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.BayesianRidgeEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func TweedieRegressorResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "tweedie regressor"

		var data pb.TweedieRegressorRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.TweedieRegressorEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func SGDClassifierResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "sgd classifier"

		var data pb.SGDClassifierRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.SGDClassifierEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}

func ElasticNetResponse(linearClient *LinearGRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		const event_name = "elastic net"

		var data pb.ElasticNetRequest
		if err := ReadRequestData(c, &data); err != nil {
			HandleBadRequestError(c, err, event_name)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := linearClient.client.ElasticNetEvent(ctx, &data)

		if err != nil {
			HandleGRPCError(c, err, event_name)
			return
		}

		SendResponse(c, response, event_name)

	}

}
